package cmd

import (
	"fmt"

	"github.com/digitallysavvy/ten-agent-cli/internal/deployer"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the TEN Agent",
	Long: `Package and deploy the TEN Agent to a specified environment.

This command prepares your TEN Agent project for deployment and
sends it to the configured deployment environment.

Example:
  ten-agent deploy`,
	Run: func(cmd *cobra.Command, args []string) {
		err := deployer.Deploy()
		if err != nil {
			fmt.Printf("Deployment failed: %v\n", err)
			return
		}
		fmt.Println("Deployment completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	// Add flags for deployment options
}
