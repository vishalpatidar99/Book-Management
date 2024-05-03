package main

import (
	"github.com/vishalpatidar99/Book-Management/api"
	"github.com/vishalpatidar99/Book-Management/db"
)

func main() {
	db.ConnectDatabase()

	router := api.SetupRouter()
	router.Run(":8080")
}
