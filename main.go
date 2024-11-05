package main

import (
	"log"
	"net/http"

	"github.com/sanjeevmalagi1/go-crud-api/db"
	"github.com/sanjeevmalagi1/go-crud-api/handlers"
)

func main() {
	db.InitDB("postgres://postgres:postgres@localhost:5432/crud_db?sslmode=disable")

	http.HandleFunc("/user", handlers.CreateUser)
	http.HandleFunc("/users", handlers.GetUsers)

	log.Println("Start server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
