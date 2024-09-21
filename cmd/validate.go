package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/digitallysavvy/ten-agent-cli/internal/validator"
)

var validateCmd = &cobra.Command{
    Use:   "validate",
    Short: "Validate the TEN Agent project",
    Long:  `Check for common configuration errors and validate graph connections.`,
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