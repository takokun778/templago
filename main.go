package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		_, _ = w.Write([]byte(Hello("Golang")))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT 1")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			var v int
			rows.Scan(&v)
			log.Printf("Health: %d", v)
		}

		_, _ = w.Write([]byte("OK"))
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

	log.Printf("server listen on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
