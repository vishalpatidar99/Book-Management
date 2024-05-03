package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"user_type"` // Admin or Regular
}

type Book struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	User            string `json:"user"`
	PublicationYear int    `json:"publish_year"`
}
