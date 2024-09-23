package cmd

import (
	"fmt"

	"github.com/digitallysavvy/ten-agent-cli/internal/validator"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the TEN Agent project",
	Long: `Check for common configuration errors and validate graph connections.

This command performs a series of checks on your TEN Agent project to ensure
that all configurations are valid and the graph connections are correct.

Example:
  ten-agent validate`,
	Run: func(cmd *cobra.Command, args []string) {
		err := validator.Validate()
		if err != nil {
			fmt.Printf("Validation failed: %v\n", err)
			return
		}
		fmt.Println("Validation completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
