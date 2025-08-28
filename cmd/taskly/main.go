package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/erfuuan/taskly-go/internal/db"
	"github.com/erfuuan/taskly-go/internal/server"
	"github.com/erfuuan/taskly-go/internal/service"
	"github.com/erfuuan/taskly-go/internal/tasks"
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

	// Initialize tasks table
	if err := tasks.InitTasksTable(); err != nil {
		fmt.Println("❌ Error initializing tasks table:", err)
		return
	}

	// REST API server
	go func() {
		fmt.Println("🚀 REST API server starting on port 3000...")
		server.Start()
	}()

	// Task runner
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
}
