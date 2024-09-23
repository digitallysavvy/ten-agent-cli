// cmd/init.go
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/digitallysavvy/ten-agent-cli/internal/project"
	"github.com/spf13/cobra"
)

var (
	agoraAppID        string
	agoraCertificate  string
	azureSTTKey       string
	azureSTTRegion    string
	azureTTSKey       string
	azureTTSRegion    string
	openAIKey         string
	graphDesignerPort string
	serverPort        string
	workersMax        string
	workerQuitTimeout string
	logPath           string
	logStdout         string
	verbose           bool
)

var initCmd = &cobra.Command{
	Use:   "init [project name]",
	Short: "Initialize a new TEN Agent project",
	Long: `Initialize a new TEN Agent project using the TEN-Agent template from GitHub.
This command will set up the project structure and install required components.

Example:
  ten-agent init myproject --agora-app-id=YOUR_AGORA_APP_ID --openai-key=YOUR_OPENAI_KEY`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		if verbose {
			log.Printf("Initializing project: %s\n", projectName)
		}

		envVars := map[string]string{
			"AGORA_APP_ID":                 getOrPrompt(agoraAppID, "Enter Agora App ID", true),
			"AGORA_APP_CERTIFICATE":        getOrPrompt(agoraCertificate, "Enter Agora App Certificate", false),
			"AZURE_STT_KEY":                getOrPrompt(azureSTTKey, "Enter Azure Speech-to-Text Key", true),
			"AZURE_STT_REGION":             getOrPrompt(azureSTTRegion, "Enter Azure Speech-to-Text Region", true),
			"AZURE_TTS_KEY":                getOrPrompt(azureTTSKey, "Enter Azure Text-to-Speech Key", true),
			"AZURE_TTS_REGION":             getOrPrompt(azureTTSRegion, "Enter Azure Text-to-Speech Region", true),
			"OPENAI_API_KEY":               getOrPrompt(openAIKey, "Enter OpenAI API Key", true),
			"GRAPH_DESIGNER_SERVER_PORT":   getOrPrompt(graphDesignerPort, "Enter Graph Designer Server Port", false),
			"SERVER_PORT":                  getOrPrompt(serverPort, "Enter Server Port", false),
			"WORKERS_MAX":                  getOrPrompt(workersMax, "Enter Maximum number of workers", false),
			"WORKER_QUIT_TIMEOUT_SECONDES": getOrPrompt(workerQuitTimeout, "Enter Worker quit timeout in seconds", false),
			"LOG_PATH":                     getOrPrompt(logPath, "Enter Log path", false),
			"LOG_STDOUT":                   getOrPrompt(logStdout, "Log to stdout? (true/false)", false),
		}

		if verbose {
			log.Println("Environment variables collected:")
			for k, v := range envVars {
				if k == "OPENAI_API_KEY" || k == "AZURE_STT_KEY" || k == "AZURE_TTS_KEY" {
					log.Printf("  %s: ********\n", k)
				} else {
					log.Printf("  %s: %s\n", k, v)
				}
			}
		}

		err := project.Initialize(projectName, envVars)
		if err != nil {
			log.Fatalf("Failed to initialize project: %v\n", err)
		}

		if verbose {
			log.Printf("Project initialized successfully: %s\n", projectName)
			log.Printf("Project directory: %s\n", filepath.Join(".", projectName))
			log.Println("Files and directories created:")
			err := project.ListCreatedFiles(projectName)
			if err != nil {
				log.Printf("Error listing created files: %v\n", err)
			}
		}

		fmt.Printf("Successfully initialized project: %s\n", projectName)
		fmt.Println("Next steps:")
		fmt.Println("1. cd", projectName)
		fmt.Println("2. Review the .env file to ensure all environment variables are set correctly")
		fmt.Println("3. Run 'ten-agent start' to start the TEN Agent services")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVar(&agoraAppID, "agora-app-id", "", "Agora App ID (required)")
	initCmd.Flags().StringVar(&agoraCertificate, "agora-certificate", "", "Agora App Certificate")
	initCmd.Flags().StringVar(&azureSTTKey, "azure-stt-key", "", "Azure Speech-to-Text Key (required)")
	initCmd.Flags().StringVar(&azureSTTRegion, "azure-stt-region", "", "Azure Speech-to-Text Region (required)")
	initCmd.Flags().StringVar(&azureTTSKey, "azure-tts-key", "", "Azure Text-to-Speech Key (required)")
	initCmd.Flags().StringVar(&azureTTSRegion, "azure-tts-region", "", "Azure Text-to-Speech Region (required)")
	initCmd.Flags().StringVar(&openAIKey, "openai-key", "", "OpenAI API Key (required)")
	initCmd.Flags().StringVar(&graphDesignerPort, "graph-designer-port", "49483", "Graph Designer Server Port")
	initCmd.Flags().StringVar(&serverPort, "server-port", "8080", "Server Port")
	initCmd.Flags().StringVar(&workersMax, "workers-max", "10", "Maximum number of workers")
	initCmd.Flags().StringVar(&workerQuitTimeout, "worker-quit-timeout", "60", "Worker quit timeout in seconds")
	initCmd.Flags().StringVar(&logPath, "log-path", "/tmp/astra", "Log path")
	initCmd.Flags().StringVar(&logStdout, "log-stdout", "false", "Log to stdout (true/false)")
	initCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
}

func getOrPrompt(value, prompt string, required bool) string {
	if value != "" {
		return value
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		if required {
			fmt.Print(prompt + " (required): ")
		} else {
			fmt.Print(prompt + " (optional, press Enter to skip): ")
		}
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" || !required {
			return input
		}
		fmt.Println("This field is required. Please enter a value.")
	}
}
