package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "Get help on any command",
	Long: `Help provides help for any command in the application.
Simply type ten-agent help [path to command] for full details.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("TEN Agent CLI - Available Commands:")
			fmt.Println()
			for _, command := range rootCmd.Commands() {
				fmt.Printf("%-15s %s\n", command.Use, command.Short)
			}
			fmt.Println()
			fmt.Println("Use 'ten-agent help [command]' for more information about a command.")
		} else {
			targetCmd, _, err := rootCmd.Find(args)
			if err != nil {
				fmt.Printf("Unknown help topic %#q\n", args)
				cmd.Root().Usage()
			} else {
				targetCmd.HelpFunc()(targetCmd, args)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
