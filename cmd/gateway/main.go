package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("POLARIS_PORT")
	if port == "" { port = "8800" }
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "up", "service": "gateway"})
	})
	log.Printf("polaris gateway on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
