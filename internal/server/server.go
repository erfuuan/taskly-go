package server

import (
	"fmt"
	"net/http"
	"os"
)

func Start() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	fmt.Printf("🚀 Taskly REST API running on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("❌ REST API server error:", err)
	}
}
