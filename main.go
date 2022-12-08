package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, Golang!"))
	})

	log.Print("server run...")

	const timeout = 3

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: timeout * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
