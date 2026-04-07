package main

import (
	"fmt"
	"go-backend/db"
	"go-backend/routes"
	"net/http"
)

func main() {
	db.Connect()

	routes.RegisterRoutes()

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}