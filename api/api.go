package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalpatidar99/Book-Management/db"
	"github.com/vishalpatidar99/Book-Management/models"
)

func LoginAPI(c *gin.Context) {
	var user models.User
	err := c.BindJSON(user)
	if err != nil || (user.Type != "Admin" && user.Type != "Regular") {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request data"})
		return
	}

	if err := db.DB.Create(&user); err != nil {
		panic(err)
	}

	c.Status(http.StatusCreated)
}

func UserHomeAPI(c *gin.Context) {

}

func AddBookAPI(c *gin.Context) {

}

func DeleteBookAPI(c *gin.Context) {

}
