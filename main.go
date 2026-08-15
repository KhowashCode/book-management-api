package main

import (
	"book-management-api/database"
	"book-management-api/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	database.ConnectToDatabase()
	routes.Server().Run(":" + port)
}
