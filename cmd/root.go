/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "password-generator",
	Short: "A simple CLI tool to generate secure passwords",
	Long: `password-generator is a command line tool that helps you generate secure and random passwords for your applications and services. 
It offers various options to customize the length, complexity, and character sets of the generated passwords.

Features:
- Generate passwords of any length
- Include or exclude uppercase letters, lowercase letters, numbers, and special characters
- Ensure password complexity to meet security standards
- Generate multiple passwords at once
- Copy the generated password to the clipboard

Examples:
- Generate a simple 12-character password:
	password-generator -l 12

- Generate a complex 16-character password with special characters:
	password-generator -l 16 -s

- Generate 5 passwords of 20 characters each:
	password-generator -l 20 -n 5

This tool is built using the Cobra library for Go, which provides a simple and powerful way to create command line applications.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.password-gen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
