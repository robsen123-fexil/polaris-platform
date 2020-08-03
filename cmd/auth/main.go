package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("POLARIS_PORT")
	if port == "" { port = "8801" }
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "up", "service": "auth"})
	})
	log.Printf("polaris auth on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
