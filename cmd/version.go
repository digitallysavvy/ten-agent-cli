package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.11" // TODO: use build flags to set version

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TEN Agent CLI",
	Long:  `All software has versions. This is TEN Agent's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("TEN Agent CLI v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
