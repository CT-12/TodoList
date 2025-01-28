/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"todo/internal/model"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [Task name]",
	Short: "Delete a task from todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := model.Connect()
		if err != nil {
			log.Fatalln("failed to connect database")
			os.Exit(1)
		}
		defer model.CloseDB(db)

		taskName := args[0]
		model.DeleteTaskByName(db, taskName)

		log.Println("Task deleted successfully")

		showTasks(model.GetAll(db))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
