package extension

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	anthropicAPIURL = "https://api.anthropic.com/v1/messages"
	anthropicModel  = "claude-3-sonnet-20240229"
)

var anthropicAPIKey string

type AnthropicRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AnthropicResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

func Generate(name string, buildTimeAnthropicAPIKey string, verbose bool) error {
	if verbose {
		log.Println("Starting AI-powered extension generation...")
	}

	if buildTimeAnthropicAPIKey == "" {
		return fmt.Errorf("ANTHROPIC_API_KEY build-time variable not set")
	}

	// Set the anthropicAPIKey to the build-time variable
	anthropicAPIKey = buildTimeAnthropicAPIKey

	// Create the extension using the Create function
	err := Create(name)
	if err != nil {
		return fmt.Errorf("failed to create extension: %w", err)
	}

	extensionDescription := askQuestion("Briefly describe what your extension does: ")
	apiKey := askQuestion("Enter the name of the API key property (e.g., anthropic_api_key): ")
	apiURL := askQuestion("Enter the API URL: ")
	modelName := askQuestion("Enter the model name to use: ")

	if verbose {
		log.Println("Reading supplemental files...")
	}
	supplementalContent, err := readSupplementalFiles()
	if err != nil {
		return fmt.Errorf("failed to read supplemental files: %w", err)
	}

	if verbose {
		log.Println("Generating extension files using Claude 3.5 Sonnet...")
	}

	basePrompt := fmt.Sprintf(`You are an AI assistant specialized in creating TEN Framework extensions. Use the following supplemental information as reference, but adapt it for the specific extension being created. 

Supplemental Information:
%s

Now, generate the content for a TEN Framework extension with the following details:
- Name: %s
- Description: %s
- API Key: %s
- API URL: %s
- Model Name: %s

The extension should implement the TEN Framework interfaces and handle functions like data_in, data_out, and cmd_in, cmd_out, audio_in, and audio_out, video_in, and video_out. 

IMPORTANT: Provide ONLY the code without any explanations or comments. Start your response with the opening code fence and end with the closing code fence. Make sure to use 'ten_framework/ten' instead of 'agoraio/rte' in the import statements.
`, supplementalContent, name, extensionDescription, apiKey, apiURL, modelName)

	goCodePrompt := basePrompt + "\nGenerate the main.go file for this extension:"
	manifestJSONPrompt := basePrompt + "\nGenerate the manifest.json file for this extension:"
	propertyJSONPrompt := basePrompt + "\nGenerate the property.json file for this extension:"

	goCode, err := generateWithClaude(goCodePrompt, verbose)
	if err != nil {
		return fmt.Errorf("failed to generate main.go: %w", err)
	}
	goCode = sanitizeResponse(goCode)

	manifestJSON, err := generateWithClaude(manifestJSONPrompt, verbose)
	if err != nil {
		return fmt.Errorf("failed to generate manifest.json: %w", err)
	}
	manifestJSON = sanitizeResponse(manifestJSON)

	propertyJSON, err := generateWithClaude(propertyJSONPrompt, verbose)
	if err != nil {
		return fmt.Errorf("failed to generate property.json: %w", err)
	}
	propertyJSON = sanitizeResponse(propertyJSON)

	// Define the path for the new extension
	extensionPath := filepath.Join("ten_packages", "extension", name)

	//Overwrite the files with the generated code
	err = os.WriteFile(filepath.Join(extensionPath, name+"_extension.go"), []byte(goCode), 0644)
	if err != nil {
		return fmt.Errorf("failed to write %s_extension.go: %w", name, err)
	}

	err = os.WriteFile(filepath.Join(extensionPath, "manifest.json"), []byte(manifestJSON), 0644)
	if err != nil {
		return fmt.Errorf("failed to write manifest.json: %w", err)
	}

	err = os.WriteFile(filepath.Join(extensionPath, "property.json"), []byte(propertyJSON), 0644)
	if err != nil {
		return fmt.Errorf("failed to write property.json: %w", err)
	}

	if verbose {
		log.Printf("Extension files have been generated and copied into the '%s' directory.\n", extensionPath)
	}

	return nil
}

func askQuestion(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	answer, _ := reader.ReadString('\n')
	return strings.TrimSpace(answer)
}

func generateWithClaude(prompt string, verbose bool) (string, error) {
	if verbose {
		log.Println("Preparing request to Claude API...")
	}

	request := AnthropicRequest{
		Model:     anthropicModel,
		MaxTokens: 4000,
		Messages:  []Message{{Role: "user", Content: prompt}},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	if verbose {
		log.Println("Sending request to Claude API...")
	}

	req, err := http.NewRequest("POST", anthropicAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", anthropicAPIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if verbose {
		log.Printf("Received response with status code: %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if verbose {
		log.Printf("Response body: %s\n", string(body))
	}

	// Check if the response status code indicates an error
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned non-200 status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var response AnthropicResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(response.Content) == 0 {
		return "", fmt.Errorf("empty response content from API")
	}

	return response.Content[0].Text, nil
}

func readSupplementalFiles() (string, error) {
	baseURL := "https://raw.githubusercontent.com/digitallysavvy/tenframework-docs/main/"
	files := []string{
		"getting_started/quickstart.md",
		"getting_started/create_a_hello_world_extension.md",
		"ten_service/ten_api_beta.md",
		"ten_service/ten_schema_beta.md",
		"tutorials/how_to_build_extension_with_go_beta.md",
		"ten_framework/api/required.md",
		"ten_framework/building.md",
		"ten_framework/concept_overview.md",
		"ten_framework/dependencies.md",
		"manifest.json",
		"property.json",
		"main.go",
	}

	var supplementalContent strings.Builder
	for _, file := range files {
		url := baseURL + file
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Warning: Couldn't fetch file %s. Error: %v\n", file, err)
			continue
		}
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Warning: Couldn't read content from %s. Error: %v\n", file, err)
			continue
		}
		supplementalContent.WriteString(fmt.Sprintf("File: %s\n\n%s\n\n", file, string(content)))
	}
	return supplementalContent.String(), nil
}

// sanitizeResponse removes the code fences and any leading newlines from the response
func sanitizeResponse(response string) string {
	// Find the first occurrence of ```
	start := strings.Index(response, "```")
	if start == -1 {
		return response
	}
	// Find the closing ```
	end := strings.LastIndex(response, "```")
	if end == -1 || end <= start {
		return response
	}
	// Extract the content between the code fences
	content := response[start+3 : end]
	// Remove the language identifier and any leading newlines
	lines := strings.SplitN(content, "\n", 2)
	if len(lines) > 1 {
		content = strings.TrimSpace(lines[1])
	} else {
		content = strings.TrimSpace(content)
	}
	return content
}
