package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/erfuuan/taskly-go/internal/db"
	"github.com/erfuuan/taskly-go/internal/server"
	"github.com/erfuuan/taskly-go/internal/tasks"
)

func main() {
	godotenv.Load()

	if err := db.Connect(); err != nil {
		fmt.Println("❌ Failed to connect to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	var rootCmd = &cobra.Command{
		Use:   "taskly",
		Short: "Taskly CLI - Manage your tasks",
	}

	// Add command
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new task",
		Run: func(cmd *cobra.Command, args []string) {
			title, _ := cmd.Flags().GetString("title")
			duedateStr, _ := cmd.Flags().GetString("duedate")
			dueDate, _ := time.Parse("2006-01-02-15:04", duedateStr)

			if err := tasks.AddTask(title, dueDate); err != nil {
				fmt.Println("❌ Error:", err)
			} else {
				fmt.Println("✅ Task added successfully:", title)
			}
		},
	}
	addCmd.Flags().String("title", "", "Task title")
	addCmd.Flags().String("duedate", "", "Due date in YYYY-MM-DD-HH:mm")
	addCmd.MarkFlagRequired("title")
	addCmd.MarkFlagRequired("duedate")

	// Run command
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Check tasks due in 30 minutes",
		Run: func(cmd *cobra.Command, args []string) {
			msgs, err := tasks.RunTasks()
			if err != nil {
				fmt.Println("❌ Error:", err)
				return
			}
			if len(msgs) == 0 {
				fmt.Println("⚠️  No tasks due in the next 30 minutes.")
			}
			for _, m := range msgs {
				fmt.Println("⏰", m)
			}
		},
	}

	// REST API command
	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start REST API",
		Run: func(cmd *cobra.Command, args []string) {
			server.Start()
		},
	}

	rootCmd.AddCommand(addCmd, runCmd, serveCmd)
	rootCmd.Execute()
}
