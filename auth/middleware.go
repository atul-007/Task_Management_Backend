// auth.go

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the authentication token exists in the header
		tokenString, err := c.Cookie("Authorization")
		if tokenString == "" || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
			c.Abort()
			return
		}

		// Verify the token
		userID, err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
			c.Abort()
			return
		}

		// Set the user ID in the context for later use if needed
		c.Set("userID", userID)

		// If the token is valid, continue with the request
		c.Next()
	}
}
