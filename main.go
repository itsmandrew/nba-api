package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Configured Go Web Server!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server starting with custom configurations on port 8080...")
	log.Fatal(server.ListenAndServe())
}
