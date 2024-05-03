package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"user_type"` // Admin or Regular
}

type Book struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publish_year"`
}
