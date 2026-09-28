package main

import (
	"log"
	"product-api/db"
)

func main() {

	database, err := db.Connect()

	if err != nil {
		log.Fatalf("Database connection failed to connect")
	}

	defer database.Close()

	log.Println("Database connected succesfully")

	if err := startServer(); err != nil {
		log.Fatalf("Error starting server %v", err)
	}

}
