package project

import (
	"fmt"
	"os"
	"os/exec"
)

func Initialize(projectName string, envVars map[string]string) error {
	// Create project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Change to project directory
	err = os.Chdir(projectName)
	if err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}

	// Clone the TEN-Agent template directly into the current directory
	cmd := exec.Command("git", "clone", "--depth", "1", "https://github.com/TEN-framework/TEN-Agent", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone TEN-Agent template: %w\n%s", err, output)
	}

	// Remove the .git directory to disassociate from the template repository
	err = os.RemoveAll(".git")
	if err != nil {
		return fmt.Errorf("failed to remove .git directory: %w", err)
	}

	// Initialize a new git repository
	cmd = exec.Command("git", "init")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize new git repository: %w", err)
	}

	// Create .env file with provided environment variables
	err = createEnvFile(envVars)
	if err != nil {
		return fmt.Errorf("failed to create .env file: %w", err)
	}

	// Install dependencies
	cmd = exec.Command("docker", "compose", "pull")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to pull Docker images: %w", err)
	}

	// Build the project
	cmd = exec.Command("docker", "compose", "run", "--rm", "astra_agents_dev", "make", "build")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to build the project: %w", err)
	}

	return nil
}

func createEnvFile(envVars map[string]string) error {
	file, err := os.Create(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	for key, value := range envVars {
		if value != "" {
			_, err := file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
