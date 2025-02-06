package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Record struct definition
type Record struct {
	ID       int               `json:"id"`
	Values   []float64         `json:"values"`
	Metadata map[string]string `json:"metadata"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("Processed")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Processed"))
}

func main() {
	http.HandleFunc("/upload", handler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
