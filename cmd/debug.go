package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/digitallysavvy/ten-agent-cli/internal/debug"
)

var debugCmd = &cobra.Command{
    Use:   "debug",
    Short: "Debug utilities for TEN Agent",
    Long:  `Provides debugging information and utilities for TEN Agent.`,
}

var logsCmd = &cobra.Command{
    Use:   "logs",
    Short: "View Agent logs",
    Run: func(cmd *cobra.Command, args []string) {
        err := debug.ViewLogs()
        if err != nil {
            fmt.Printf("Failed to view logs: %v\n", err)
            return
        }
    },
}

func init() {
    rootCmd.AddCommand(debugCmd)
    debugCmd.AddCommand(logsCmd)
}