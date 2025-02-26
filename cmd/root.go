/*
Copyright © 2025 Chia-Tai, Lee <ctlee.8912@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	VERSION string = "0.1.0"
)

var (
	show_version bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "This CLI tool is a simple todo list application.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		if show_version {
			cmd.Println("Version: ", VERSION)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.To-do-list.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&show_version, "version", "v", false, "Show version")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}




