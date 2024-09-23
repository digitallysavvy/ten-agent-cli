package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ten-agent",
	Short: "TEN Agent CLI - A tool for building and managing TEN framework Agents",
	Long: `TEN Agent CLI is a comprehensive command-line tool designed to help developers 
build, manage, and deploy AI Agents using the TEN framework.

It provides commands for initializing projects, managing configurations, 
creating extensions, running and debugging agents, and more.`,
	Example: `  ten-agent init myproject
  ten-agent run
  ten-agent config edit
    Use: 'ten-agent help' to see all available commands.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define flags and configuration settings for the root command
}
