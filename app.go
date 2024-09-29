package main

import (
	"log"

	"yosbomb.com/bucketbudget/server"
)

func main() {
	log.Println("Starting server!")

	server := server.NewServer()

	log.Fatal(server.ListenAndServe(":8080"))

}
