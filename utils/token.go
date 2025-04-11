package utils

import (
	"Task_manager_apis/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims defines a struct for your JWT claims, making it more type-safe
type CustomClaims struct {
	UserID uint   `json:"sub"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID uint, email string) (string, error) {
	// Set JWT claims
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			// Expiration time (24 hours)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	secretKey := config.GetSecretKey()
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims if valid
func ValidateToken(tokenString string) (CustomClaims, error) {
	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		// Return the secret key to validate the JWT token
		return []byte(config.GetSecretKey()), nil
	})

	if err != nil || !token.Valid {
		return CustomClaims{}, errors.New("invalid or expired token")
	}

	// Return the claims if valid
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return *claims, nil
	}

	return CustomClaims{}, errors.New("invalid token claims")
}
