package cmd

import (
	"fmt"
	"log"

	"github.com/digitallysavvy/ten-agent-cli/internal/extension"
	"github.com/spf13/cobra"
)

var (
	extensionName string
)

var generateExtensionCmd = &cobra.Command{
	Use:   "generate-extension [extension name]",
	Short: "Generate a new TEN Agent extension",
	Long: `Generate a new TEN Agent extension using a template.
This command will set up the extension structure and create required files.

Example:
  ten-agent generate-extension myextension`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		extensionName = args[0]
		if verbose {
			log.Printf("Generating extension: %s\n", extensionName)
		}

		err := extension.Generate(extensionName, verbose)
		if err != nil {
			log.Fatalf("Failed to generate extension: %v\n", err)
		}

		if verbose {
			log.Printf("Extension generated successfully: %s\n", extensionName)
		}

		fmt.Printf("Successfully generated extension: %s\n", extensionName)
		fmt.Println("Next steps:")
		fmt.Println("1. Review the generated files in the 'agents' directory")
		fmt.Println("2. Implement your extension logic in the generated files")
		fmt.Println("3. Update the manifest.json file if needed")
	},
}

func init() {
	rootCmd.AddCommand(generateExtensionCmd)
	generateExtensionCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
}
