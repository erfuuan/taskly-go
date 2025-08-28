package server

import (
	"fmt"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"ok"}`)
	})

	fmt.Println("🚀 Taskly REST API running on :3000")
	http.ListenAndServe(":3000", mux)
}
