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
	"github.com/erfuuan/taskly-go/service"
)

func waitForDB(maxRetries int, delay time.Duration) {
	for i := 0; i < maxRetries; i++ {
		if err := db.Connect(); err != nil {
			fmt.Println("Waiting for database...", err)
			time.Sleep(delay)
		} else {
			fmt.Println("✅ Connected to database")
			return
		}
	}
	fmt.Println("❌ Could not connect to database after retries")
	os.Exit(1)
}

func main() {
	godotenv.Load()
	waitForDB(10, 3*time.Second)
	defer db.Close()

	if err := tasks.InitTasksTable(); err != nil {
		fmt.Println("❌ Error initializing tasks table:", err)
		return
	}

	rootCmd := &cobra.Command{
		Use:   "taskly",
		Short: "Taskly CLI & REST API",
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				fmt.Println("🚀 REST API server starting on port 3000...")
				server.Start()
			}()

			go func() {
				for {
					msgs, err := service.RunTasks()
					if err != nil {
						fmt.Println("❌ Task runner error:", err)
					} else if len(msgs) > 0 {
						for _, m := range msgs {
							fmt.Println("⏰", m)
						}
					}
					time.Sleep(30 * time.Second)
				}
			}()

			fmt.Println("✅ Both REST API and Task Runner are running concurrently")
			select {} // block forever
		},
	}

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new task",
		Run: func(cmd *cobra.Command, args []string) {
			title, _ := cmd.Flags().GetString("title")
			duedateStr, _ := cmd.Flags().GetString("duedate")

			dueDate, err := time.Parse("2006-01-02-15:04", duedateStr)
			if err != nil {
				fmt.Println("❌ Invalid date format. Use YYYY-MM-DD-HH:mm")
				return
			}

			if err := service.AddTask(title, dueDate); err != nil {
				fmt.Println("❌ Error adding task:", err)
			} else {
				fmt.Println("✅ Task added successfully:", title)
			}
		},
	}

	addCmd.Flags().String("title", "", "Task title")
	addCmd.Flags().String("duedate", "", "Due date in YYYY-MM-DD-HH:mm")
	addCmd.MarkFlagRequired("title")
	addCmd.MarkFlagRequired("duedate")

	rootCmd.AddCommand(addCmd)

	rootCmd.Execute()
}
