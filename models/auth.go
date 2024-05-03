package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}
