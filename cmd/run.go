package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/digitallysavvy/ten-agent-cli/internal/runtime"
)

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Run the TEN Agent",
    Long:  `Start a TEN Agent with specified configurations.`,
    Run: func(cmd *cobra.Command, args []string) {
        err := runtime.RunAgent()
        if err != nil {
            fmt.Printf("Failed to run agent: %v\n", err)
            return
        }
    },
}

func init() {
    rootCmd.AddCommand(runCmd)
    // Add flags for agent configuration
}