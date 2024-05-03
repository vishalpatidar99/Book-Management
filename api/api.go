package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vishalpatidar99/Book-Management/models"
	"github.com/vishalpatidar99/Book-Management/utils"
)

func LoginAPI(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil || (user.Type != "Admin" && user.Type != "Regular") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	token, err := LoginHandler(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(200, token)
}

func UserHomeAPI(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		return
	}

	tokenString := strings.Split(authHeader, "Bearer ")[1]
	claims, ok := utils.ParseAndValidateToken(tokenString)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	books, err := UserHomeHandler(claims.UserType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, books)
}

func AddBookAPI(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		return
	}

	tokenString := strings.Split(authHeader, "Bearer ")[1]
	claims, ok := utils.ParseAndValidateToken(tokenString)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	if claims.UserType != "Admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "only admin users can access this endpoint"})
		return
	}

	var req models.Book
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if req.Name == "" || req.Author == "" || req.PublicationYear <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
		return
	}

	if err := AddBookHandler(req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to add book in regularUser.csv"})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteBookAPI(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		return
	}

	tokenString := strings.Split(authHeader, "Bearer ")[1]

	// Parse and validate the JWT token
	claims, ok := utils.ParseAndValidateToken(tokenString)
	if !ok || claims.UserType != "Admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or unauthorized token"})
		return
	}

	// Bind request query parameter to book name
	bookName := c.Query("name")
	if bookName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "book name parameter is required"})
		return
	}

	if err := DeleteBookHandler(bookName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to delete book from regularUser.csv"})
		return
	}

	c.Status(http.StatusOK)
}
