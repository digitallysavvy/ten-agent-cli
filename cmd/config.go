package cmd

import (
	"fmt"

	"github.com/digitallysavvy/ten-agent-cli/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage TEN Agent configuration",
	Long: `Edit Agent configurations and manage environment variables.

Subcommands:
  edit    Edit the Agent configuration

Usage:
  ten-agent config [command]

Example:
  ten-agent config edit`,
}

var editConfigCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the Agent configuration",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.EditConfig()
		if err != nil {
			fmt.Printf("Failed to edit configuration: %v\n", err)
			return
		}
		fmt.Println("Configuration updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(editConfigCmd)
}
