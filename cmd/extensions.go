package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/digitallysavvy/ten-agent-cli/internal/container"
    "github.com/digitallysavvy/ten-agent-cli/internal/extension"
)

var createExtensionCmd = &cobra.Command{
    Use:   "create-extension [name]",
    Short: "Create a new custom extension",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        if !container.IsInsideContainer() {
            fmt.Println("This command must be run inside the TEN framework container.")
            return
        }
        name := args[0]
        err := extension.Create(name)
        if err != nil {
            fmt.Printf("Failed to create extension: %v\n", err)
            return
        }
        fmt.Printf("Successfully created extension: %s\n", name)
    },
}

func init() {
    rootCmd.AddCommand(createExtensionCmd)
}