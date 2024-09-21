package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/digitallysavvy/ten-agent-cli/internal/config"
)

var configCmd = &cobra.Command{
    Use:   "config",
    Short: "Manage TEN Agent configuration",
    Long:  `Edit Agent configurations and manage environment variables.`,
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