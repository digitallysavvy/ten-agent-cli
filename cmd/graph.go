package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/digitallysavvy/ten-agent-cli/internal/graph"
)

var graphCmd = &cobra.Command{
    Use:   "graph",
    Short: "Manage TEN graphs",
    Long:  `Launch the Graph Editor for creating and editing TEN graphs.`,
    Run: func(cmd *cobra.Command, args []string) {
        err := graph.LaunchEditor()
        if err != nil {
            fmt.Printf("Failed to launch Graph Editor: %v\n", err)
            return
        }
    },
}

func init() {
    rootCmd.AddCommand(graphCmd)
}