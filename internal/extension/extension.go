package extension

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

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
