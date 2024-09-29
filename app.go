package main

import (
	"log"
)

func main() {
	log.Println("Starting server!")

	server := NewServer()

	log.Fatal(server.ListenAndServe(":8080"))

}
