package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func Start() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp := HealthResponse{
			Status:    "ok",
			Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		}

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Failed to generate JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	})

	fmt.Printf("🚀 Taskly REST API running on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("❌ REST API server error:", err)
	}
}
