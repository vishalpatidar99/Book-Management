package api

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/vishalpatidar99/Book-Management/auth"
	"github.com/vishalpatidar99/Book-Management/db"
	"github.com/vishalpatidar99/Book-Management/models"
	"github.com/vishalpatidar99/Book-Management/utils"
)

func LoginHandler(user models.User) (string, error) {

	res := db.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
	if res.RowsAffected == 0 {
		return "", errors.New("user not found or invalid credentials")
	}

	token, err := auth.GetJWTToken(user.Username, user.Type)
	if err != nil {
		return "", err
	}

	return token, nil
}

func UserHomeHandler(userType string) ([]models.Book, error) {
	var books []models.Book
	regularBooks, err := utils.ReadCSV("regularUser.csv")
	if err != nil {
		return []models.Book{}, err
	}
	books = append(books, regularBooks...)

	if userType == "Admin" {
		// reding adminUser.csv too
		adminBooks, err := utils.ReadCSV("adminUser.csv")
		if err != nil {
			return []models.Book{}, err
		}
		books = append(books, adminBooks...)
	}

	return books, nil
}

func AddBookHandler(req models.Book) error {
	file, err := os.OpenFile("regularUser.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write([]string{req.Name, req.Author, strconv.Itoa(req.PublicationYear)})
	if err != nil {
		return err
	}

	writer.Write([]string{})

	writer.Flush()

	return nil
}

func DeleteBookHandler(bookName string) error {
	file, err := os.OpenFile("regularUser.csv", os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	var updatedRecords [][]string
	found := false
	for _, record := range records {
		if strings.EqualFold(record[0], bookName) {
			found = true
			continue
		}
		updatedRecords = append(updatedRecords, record)
	}

	if !found {
		return errors.New("book not found in regularUser.csv")
	}

	if err := file.Truncate(0); err != nil {
		return err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	if err := writer.WriteAll(updatedRecords); err != nil {
		return err
	}
	writer.Flush()

	return nil
}
