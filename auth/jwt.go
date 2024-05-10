// jwt.go

package auth

import (
	"time"

	"github.com/atul-007/task_management_backend/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey = []byte("your-secret-key")

// Claims struct to encode token claims
type Claims struct {
	UserID primitive.ObjectID `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for the provided user
func GenerateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	// Create token claims
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   user.Username,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token and returns the user ID
func VerifyToken(tokenString string) (primitive.ObjectID, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Validate token claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return primitive.NilObjectID, err
	}

	return claims.UserID, nil
}
