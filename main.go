package main

import (
	"github.com/atul-007/task_management_backend/database"
	"github.com/atul-007/task_management_backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db *mongo.Database
)

func main() {
	// Initialize MongoDB connection and set the db variable
	gin.SetMode(gin.ReleaseMode)

	database.ConnectDB()
	db = database.GetDB()

	// Initialize Gin router
	router := gin.Default()
	corsHandler := cors.Default()
	corsMiddleware := func(c *gin.Context) {
		corsHandler.HandlerFunc(c.Writer, c.Request)
	}
	router.Use(corsMiddleware)

	// Set up routes
	routes.InitializeRoutes(router, db)

	// Run the server
	port := ":8080" // Change port number as needed
	router.Run(port)
}
