package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, Golang!"))
	})

	log.Print("server run...")

	const timeout = 3

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		ReadHeaderTimeout: timeout * time.Second,
	}

	log.Printf("server listen on %s", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
