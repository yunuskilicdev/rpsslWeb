package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	FillChoices()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
