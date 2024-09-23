package cmd

import (
	"fmt"

	"github.com/digitallysavvy/ten-agent-cli/internal/runtime"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the TEN Agent",
	Long: `Start a TEN Agent with specified configurations.

This command launches the TEN Agent using the current project configuration.
Make sure you have initialized the project and set up all required environment variables before running this command.

Example:
  ten-agent run`,
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
