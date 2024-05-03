package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vishalpatidar99/Book-Management/models"
	"github.com/vishalpatidar99/Book-Management/utils"
)

var jwtKey = []byte(utils.GenerateRandomKey(32))

func GetJWTToken(username, userType string) (string, error) {
	return generateToken(username, userType)
}

func generateToken(username, userType string) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create JWT claims
	claims := &models.Claims{
		Username: username,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	return token.SignedString(jwtKey)
}
