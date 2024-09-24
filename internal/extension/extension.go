package extension

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Create(name string) error {
	reader := bufio.NewReader(os.Stdin)

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

	// Define the path for the new extension
	extensionPath := filepath.Join("ten_packages", "extension", name)

	// Rename default_extension.go to [name]_extension.go
	defaultExtensionPath := filepath.Join(extensionPath, "default_extension.go")
	mainGoPath := filepath.Join(extensionPath, name+"_extension.go")
	err = os.Rename(defaultExtensionPath, mainGoPath)
	if err != nil {
		return fmt.Errorf("failed to rename default_extension.go to main.go: %w", err)
	}

	fmt.Printf("Extension '%s' created successfully\n", name)

	fmt.Printf("Do you want to walk through the manifest options? (y/n): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	// if response is y, then walk through the manifest options
	if strings.ToLower(response) == "y" {
		// Walk through manifest.json options
		manifestPath := filepath.Join("ten_packages", "extension", name, "manifest.json")
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
	}
	return nil
}
