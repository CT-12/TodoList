/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"todo/internal/model"

	"github.com/spf13/cobra"
)

var (
	taskName string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := model.Connect()
		if err != nil {
			log.Fatal("failed to connect database")
		}
		defer model.CloseDB(db)

		model.Insert(db, &model.Task{Name: taskName, Status: "Pending"})

		log.Println("Task added successfully")

		showTasks(model.GetAll(db))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	addCmd.Flags().StringVarP(&taskName, "name", "n", "", "Task name")
	addCmd.MarkFlagRequired("name")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
