package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vishalpatidar99/Book-Management/models"
)

// GenerateRandomKey generates a random key of the specified length
func GenerateRandomKey(length int) string {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("Error generating random key: ", err)
	}
	return base64.URLEncoding.EncodeToString(key)
}

func ReadCSV(filename string) ([]models.Book, error) {
	var books []models.Book

	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Skip the first row (header row)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	// Read the CSV records
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Validate and parse the publication year
		publicationYearStr := record[2]
		if publicationYearStr == "" {
			return nil, errors.New("publication year is empty")
		}
		publicationYear, err := strconv.Atoi(publicationYearStr)
		if err != nil {
			return nil, fmt.Errorf("invalid publication year format: %v", err)
		}

		// Create a Book object and append to the list
		book := models.Book{
			Name:            record[0],
			Author:          record[1],
			PublicationYear: publicationYear,
		}
		books = append(books, book)
	}

	return books, nil
}

func ParseAndValidateToken(tokenString string) (*models.Claims, bool) {

	// Parse and validate the JWT token
	claims := &models.Claims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	return claims, true
}

// func ParseAndValidateToken(tokenString string) (*models.Claims, bool) {
// 	// Parse and validate the JWT token
// 	claims := &models.Claims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return token, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil, false
// 	}

// 	return claims, true
// }
