package cmd

import (
	"fmt"

	"github.com/digitallysavvy/ten-agent-cli/internal/graph"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Manage TEN graphs",
	Long: `Launch the Graph Editor for creating and editing TEN graphs.

This command opens the Graph Editor interface, allowing you to visually
create and modify the graph structure of your TEN Agent.

Example:
  ten-agent graph`,
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
