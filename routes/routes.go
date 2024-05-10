// routes.go

package routes

import (
	"github.com/atul-007/task_management_backend/auth"
	"github.com/atul-007/task_management_backend/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeRoutes(router *gin.Engine, db *mongo.Database) {
	// Set up routes
	api := router.Group("/api")
	{
		api.POST("/register", func(c *gin.Context) { handlers.RegisterHandler(c, db) })
		api.POST("/login", func(c *gin.Context) { handlers.LoginHandler(c, db) })

		// Protected routes with cookie-based authentication middleware
		authorized := api.Group("/")
		authorized.Use(auth.AuthMiddleware())
		{
			authorized.GET("/tasks", func(c *gin.Context) { handlers.GetTasksHandler(c, db) })
			authorized.POST("/tasks", func(c *gin.Context) { handlers.CreateTaskHandler(c, db) })
			authorized.PUT("/tasks/:id", func(c *gin.Context) { handlers.UpdateTaskHandler(c, db) })
			authorized.DELETE("/tasks/:id", func(c *gin.Context) { handlers.DeleteTaskHandler(c, db) })

			// Logout route to clear authentication token cookie
			authorized.GET("/logout", func(c *gin.Context) { handlers.LogoutHandler(c) })
		}
	}
}
