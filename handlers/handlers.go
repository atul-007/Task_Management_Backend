// handlers.go

package handlers

import (
	"context"
	"net/http"

	"github.com/atul-007/task_management_backend/auth"
	"github.com/atul-007/task_management_backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterHandler handles user registration
func RegisterHandler(c *gin.Context, db *mongo.Database) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username already exists
	existingUser, err := GetUserByUsername(user.Username, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Insert the user into the database
	_, err = db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// LoginHandler handles user login
func LoginHandler(c *gin.Context, db *mongo.Database) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists in the database
	existingUser, _ := GetUserByUsername(user.Username, db)
	if existingUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare the hashed password with the provided password
	// (You should use a secure password hashing algorithm like bcrypt)
	// if !comparePasswords(existingUser.Password, user.Password) {
	//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	//     return
	// }

	// Generate JWT token
	token, err := auth.GenerateToken(*existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.SetCookie("Authorization", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetUserByUsername retrieves a user by username from the database
func GetUserByUsername(username string, db *mongo.Database) (*models.User, error) {
	var user models.User
	err := db.Collection("users").FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User not found
		}
		return nil, err // Return other errors
	}
	return &user, nil
}

// GetTasksHandler retrieves tasks for the authenticated user
func GetTasksHandler(c *gin.Context, db *mongo.Database) {
	// Get user ID from JWT token
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, err := auth.VerifyToken(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Query tasks for the authenticated user
	var tasks []models.Task
	cursor, err := db.Collection("tasks").Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode task"})
			return
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate cursor"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// CreateTaskHandler creates a new task for the authenticated user
func CreateTaskHandler(c *gin.Context, db *mongo.Database) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from JWT token
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDStr := cookie
	userID, err := auth.VerifyToken(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	task.UserID = userID

	// Insert the task into the database
	_, err = db.Collection("tasks").InsertOne(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

// UpdateTaskHandler updates an existing task for the authenticated user
func UpdateTaskHandler(c *gin.Context, db *mongo.Database) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := auth.VerifyToken(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	task.UserID = userID

	// Update the task in the database
	taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	_, err = db.Collection("tasks").UpdateOne(context.Background(), bson.M{"_id": taskID, "user_id": userID}, bson.M{"$set": task})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTaskHandler deletes an existing task for the authenticated user
func DeleteTaskHandler(c *gin.Context, db *mongo.Database) {
	// Get user ID from JWT token
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, err := auth.VerifyToken(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Delete the task from the database
	taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	_, err = db.Collection("tasks").DeleteOne(context.Background(), bson.M{"_id": taskID, "user_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
func LogoutHandler(c *gin.Context) {
	// Clear the authentication token cookie
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
