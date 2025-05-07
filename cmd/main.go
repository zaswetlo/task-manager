package main

import (
	"log"
	"net/http"
	"task-manager/internal"
)

func main() {
	router := internal.SetupRouter()
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
