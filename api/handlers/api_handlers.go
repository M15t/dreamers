package handlers

import (
	"dreamers/models"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type data map[string]interface{}

// GetDiaries is a function that returns an echo.HandlerFunc.
// It retrieves diaries from the database using the provided gorm.DB instance
// and returns a JSON response with the retrieved diaries.
func GetDiaries(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetDiaries(db))
	}
}

// CreateDiary is a function that returns an echo.HandlerFunc.
// It binds the request body to a Diary struct, creates a new diary entry in the database using the provided gorm.DB instance,
// and returns a JSON response with the status and the created diary if successful, or an error response if there was an error.
func CreateDiary(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var diary models.Diary
		c.Bind(&diary)

		err := models.CreateDiary(db, diary)
		if err != nil {
			fmt.Println("Error:", err)
			return c.JSON(http.StatusOK, data{
				"status": false,
				"err":    err,
			})
		}

		return c.JSON(http.StatusCreated, data{
			"status": true,
			"diary":  diary,
		})

		// old method and test
		//data, err := models.CreateDiary(db, diary)
		//
		//return c.JSON(http.StatusCreated, data)
	}
}

// Login is a function that returns an echo.HandlerFunc.
// It returns a JSON response with a status of true when called.
func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, data{
			"status": true,
		})
	}
}
