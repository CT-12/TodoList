/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"time"

	"To-do-list/internal/model"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := model.Connect()
		if err != nil {
			log.Fatalln("failed to connect database")
		}
		defer model.CloseDB(db)
		tasks := model.GetAll(db)

		showTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showTasks(tasks *[]model.Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	header := table.Row{"#", "Name", "Status", "Created At", "Updated At"}
	t.AppendHeader(header)

	for i, task := range *tasks {
		// 根據任務的狀態變更顏色
		name := task.Name
		status := task.Status
		
		// 根據狀態變色
		switch status {
		case "Pending":
			name = text.Colors{text.FgHiRed}.Sprint(task.Name)
			status = text.Colors{text.FgHiRed}.Sprint(task.Status)
		case "Done":
			name = text.Colors{text.FgHiGreen}.Sprint(task.Name)
			status = text.Colors{text.FgHiGreen}.Sprint(task.Status)
		}

		// 添加行
		row := table.Row{i + 1, name, status, formatTime(task.CreatedAt), formatTime(task.UpdatedAt)}
		t.AppendRow(row)
		t.AppendSeparator()
	}

	t.Render()
}


func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}