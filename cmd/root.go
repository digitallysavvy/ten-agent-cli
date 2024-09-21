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
build, manage, and deploy AI Agents using the TEN framework.`,
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