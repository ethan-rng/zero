/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zero",
	Short: "a simple cli tool to setup code-server containers on the internet hosted off your own computer",
	Long: `zero is a CLI tool that lets you spin up and expose Docker-based code-server environments directly from your own computer.

originally built to simplify hosting live coding interviews, zero removes the hassle of configuring infrastructure by generating and managing everything you need to run isolated development environments on demand. with zero, you can quickly create reproducible, containerized workspaces and securely share them over the internet — all without relying on external hosting providers.

to get started run ` + "`$zero serve`" + `
`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zero.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
