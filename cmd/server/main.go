package main

import (
	"log"
	"../../handlers"
)

func main() {
	if err := StartServer(8000, "/tree", handlers.TreeHandler); 
	err != nil {
		log.Fatalf("Cannot initialize server: %s", err)
	}
}