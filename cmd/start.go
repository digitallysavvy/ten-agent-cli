package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the TEN Agent services",
	Long:  `Start the TEN Agent services using docker-compose and enter the container.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting TEN Agent services...")

		// Run docker-compose up
		composeCmd := exec.Command("docker", "compose", "up", "-d")
		err := composeCmd.Run()
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
