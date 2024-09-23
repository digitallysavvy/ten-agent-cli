package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the TEN Agent services",
	Long: `Stop the TEN Agent services that are running in Docker containers.

This command runs 'docker-compose down' to stop and remove all running
containers associated with the TEN Agent project.

Example:
  ten-agent stop`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping TEN Agent services...")

		// Run docker-compose down
		downCmd := exec.Command("docker-compose", "down")
		err := downCmd.Run()
		if err != nil {
			fmt.Printf("Failed to stop services: %v\n", err)
			return
		}

		fmt.Println("Services stopped successfully.")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
