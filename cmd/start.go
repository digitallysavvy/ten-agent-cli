package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the TEN Agent services",
	Long: `Start the TEN Agent services using docker-compose, build the project, and enter the container.

This command will:
1. Run 'docker compose pull' to fetch the latest images
2. Run 'docker compose run --rm astra_agents_dev make build' to build the project
3. Run 'docker compose up -d' to start the services
4. Enter the 'astra_agents_dev' container with an interactive shell

Example:
  ten-agent start`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Preparing TEN Agent services...")

		// Pull Docker images
		pullCmd := exec.Command("docker", "compose", "pull")
		err := pullCmd.Run()
		if err != nil {
			fmt.Printf("Failed to pull Docker images: %v\n", err)
			return
		}

		// Check if the image exists
		checkImageCmd := exec.Command("docker", "image", "ls", "-q", "astra_agents_dev")
		output, err := checkImageCmd.Output()
		if err != nil {
			fmt.Printf("Failed to check for existing image: %v\n", err)
			return
		}

		if strings.TrimSpace(string(output)) == "" {
			fmt.Println("Image not found. Building the project...")
			// Build the project
			buildCmd := exec.Command("docker", "compose", "run", "--rm", "astra_agents_dev", "make", "build")
			err = buildCmd.Run()
			if err != nil {
				fmt.Printf("Failed to build the project: %v\n", err)
				return
			}
			fmt.Println("Project built successfully.")
		} else {
			fmt.Println("Using existing image for astra_agents_dev.")
		}

		// Run docker-compose up
		composeCmd := exec.Command("docker", "compose", "up", "-d")
		err = composeCmd.Run()
		if err != nil {
			fmt.Printf("Failed to start services: %v\n", err)
			return
		}

		fmt.Println("Services started successfully.")

		// Enter the container
		enterCmd := exec.Command("docker", "compose", "exec", "astra_agents_dev", "/bin/bash")
		enterCmd.Stdin = os.Stdin
		enterCmd.Stdout = os.Stdout
		enterCmd.Stderr = os.Stderr

		err = enterCmd.Run()
		if err != nil {
			fmt.Printf("Failed to enter the container: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
