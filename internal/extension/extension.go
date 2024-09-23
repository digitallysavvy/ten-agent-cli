package extension

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ManifestAPI struct {
	DataIn        []APIItem `json:"data_in,omitempty"`
	DataOut       []APIItem `json:"data_out,omitempty"`
	CmdIn         []APIItem `json:"cmd_in,omitempty"`
	CmdOut        []APIItem `json:"cmd_out,omitempty"`
	InterfaceIn   []APIItem `json:"interface_in,omitempty"`
	InterfaceOut  []APIItem `json:"interface_out,omitempty"`
	AudioFrameIn  []APIItem `json:"audio_frame_in,omitempty"`
	AudioFrameOut []APIItem `json:"audio_frame_out,omitempty"`
	VideoFrameIn  []APIItem `json:"video_frame_in,omitempty"`
	VideoFrameOut []APIItem `json:"video_frame_out,omitempty"`
}

type APIItem struct {
	Name     string                 `json:"name"`
	Property map[string]interface{} `json:"property,omitempty"`
}

func Create(name string) error {
	// Change the working directory to the agents folder
	err := os.Chdir("agents")
	if err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	// Define the base command
	cmd := exec.Command("tman", "install", "extension", "default_extension_go")

	// Add template mode and data arguments
	cmd.Args = append(cmd.Args, "--template-mode")
	cmd.Args = append(cmd.Args, "--template-data", fmt.Sprintf("package_name=%s", name))
	cmd.Args = append(cmd.Args, "--template-data", fmt.Sprintf("class_name_prefix=%s", capitalize(name)))

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create extension: %w\n%s", err, output)
	}

	fmt.Printf("Extension '%s' created successfully\n", name)

	// Walk through manifest.json options
	manifestPath := filepath.Join(name, "manifest.json")
	api, err := walkThroughManifestOptions()
	if err != nil {
		return fmt.Errorf("failed to walk through manifest options: %w", err)
	}

	// Update manifest.json with new API options
	err = updateManifest(manifestPath, api)
	if err != nil {
		return fmt.Errorf("failed to update manifest: %w", err)
	}

	fmt.Println("manifest.json updated successfully")
	return nil
}

func walkThroughManifestOptions() (ManifestAPI, error) {
	api := ManifestAPI{}
	reader := bufio.NewReader(os.Stdin)

	apiTypes := []string{"data_in", "data_out", "cmd_in", "cmd_out", "interface_in", "interface_out", "audio_frame_in", "audio_frame_out", "video_frame_in", "video_frame_out"}

	for _, apiType := range apiTypes {
		fmt.Printf("Do you want to add %s? (y/n): ", apiType)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if strings.ToLower(response) == "y" {
			fmt.Print("Enter the name for this API item: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			item := APIItem{Name: name}

			if apiType == "data_in" || apiType == "data_out" {
				item.Property = make(map[string]interface{})
				fmt.Print("Enter property name (or press enter to skip): ")
				propName, _ := reader.ReadString('\n')
				propName = strings.TrimSpace(propName)

				if propName != "" {
					fmt.Print("Enter property type: ")
					propType, _ := reader.ReadString('\n')
					propType = strings.TrimSpace(propType)
					item.Property[propName] = map[string]string{"type": propType}
				}
			}

			switch apiType {
			case "data_in":
				api.DataIn = append(api.DataIn, item)
			case "data_out":
				api.DataOut = append(api.DataOut, item)
			case "cmd_in":
				api.CmdIn = append(api.CmdIn, item)
			case "cmd_out":
				api.CmdOut = append(api.CmdOut, item)
			case "interface_in":
				api.InterfaceIn = append(api.InterfaceIn, item)
			case "interface_out":
				api.InterfaceOut = append(api.InterfaceOut, item)
			case "audio_frame_in":
				api.AudioFrameIn = append(api.AudioFrameIn, item)
			case "audio_frame_out":
				api.AudioFrameOut = append(api.AudioFrameOut, item)
			case "video_frame_in":
				api.VideoFrameIn = append(api.VideoFrameIn, item)
			case "video_frame_out":
				api.VideoFrameOut = append(api.VideoFrameOut, item)
			}
		}
	}

	return api, nil
}

func updateManifest(path string, api ManifestAPI) error {
	// Read existing manifest
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var manifest map[string]interface{}
	err = json.Unmarshal(content, &manifest)
	if err != nil {
		return err
	}

	// Update API section
	manifest["api"] = api

	// Write updated manifest
	updatedContent, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, updatedContent, 0644)
}

// Helper function to capitalize the first letter of a string
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
