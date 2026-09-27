package main

import (
	"log"
)

func main() {
	if err := startServer(); err != nil {
		log.Fatalf("Error starting server %v", err);
	}
}
