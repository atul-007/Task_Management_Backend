package main

import (
	"github.com/atul-007/task_management_backend/database"
	"github.com/atul-007/task_management_backend/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db *mongo.Database
)

func main() {
	// Initialize MongoDB connection and set the db variable
	database.ConnectDB()
	db = database.GetDB()

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	routes.InitializeRoutes(router, db)

	// Run the server
	router.Run(":8080")
}
