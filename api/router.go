package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	users := v1.Group("/users")
	{
		users.POST("/login", LoginAPI)
		users.GET("/home", UserHomeAPI)
	}

	onboarding := v1.Group("/books")
	{
		onboarding.POST("/addBook", AddBookAPI)
		onboarding.DELETE("/deleteBook", DeleteBookAPI)
	}

	return router
}
