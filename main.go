package main

import (
	"ecommerce-api/db"
	"ecommerce-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitDB()

	// Create a Gin router
	router := gin.Default()

	// Initialize routes
	routes.InitRoutes(router)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
