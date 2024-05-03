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

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		publicationYearStr := record[2]
		if publicationYearStr == "" {
			return nil, errors.New("publication year is empty")
		}
		publicationYear, err := strconv.Atoi(publicationYearStr)
		if err != nil {
			return nil, fmt.Errorf("invalid publication year format: %v", err)
		}

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

	claims := &models.Claims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	return claims, true
}
