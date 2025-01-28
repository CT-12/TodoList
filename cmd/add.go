/*
Copyright © 2025 Chia-Tai, Lee <ctlee.8912@gmail.com>
*/
package cmd

import (
	"log"

	"todo/internal/model"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [Task name]",
	Short: "Add a new task to todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := model.Connect()
		if err != nil {
			log.Fatal("failed to connect database")
		}
		defer model.CloseDB(db)

		taskName := args[0]

		model.Insert(db, &model.Task{Name: taskName, Status: "Pending"})

		log.Println("Task added successfully")

		showTasks(model.GetAll(db))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
