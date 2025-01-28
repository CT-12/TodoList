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
	name string
	pending bool
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task status to 'Done' in the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := model.Connect()
		if err != nil {
			log.Fatal("failed to connect database")
		}
		defer model.CloseDB(db)

		model.UpdateOneByName(db, name, pending)

		log.Println("Task updated successfully")

		showTasks(model.GetAll(db))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.
	updateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the task to update")
	updateCmd.MarkFlagRequired("name")

	updateCmd.Flags().BoolVarP(&pending, "pending", "p", false, "Set the task status to 'pending' if true, default set to 'done'")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
