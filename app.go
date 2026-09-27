package main

import (
	"fmt"
	"net/http"
)

func startServer() error {
	fmt.Println("Server starting on port 8080")
	return http.ListenAndServe(":8080", nil)
}
