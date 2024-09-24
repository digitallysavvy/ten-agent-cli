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
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of TEN Agent CLI",
	Long:    `All software has versions. This is TEN-Agent CLI's`,
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("version")
		if v {
			fmt.Printf("TEN Agent CLI v%s (commit: %s, built at: %s)\n", version, commit, date)
		} else {
			fmt.Println(cmd.UsageString())
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
