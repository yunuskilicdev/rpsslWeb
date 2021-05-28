package main

import (
	"log"
	"net/http"
)

func main() {
	FillChoices()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}