package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TEN Agent CLI",
	Long:  `All software has versions. This is TEN Agent's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("TEN Agent CLI v%s (commit: %s, built at: %s)\n", version, commit, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
