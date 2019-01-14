package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"sample_echo/models"
	"fmt"
)

type data map[string]interface{}

func GetDiaries(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetDiaries(db))
	}
}

func CreateDiary(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var diary models.Diary
		c.Bind(&diary)

		err := models.CreateDiary(db, diary)
		if err != nil {
			fmt.Println("Error:", err)
			return c.JSON(http.StatusOK, data{
				"status": false,
				"err": err,
			})
		}

		return c.JSON(http.StatusCreated, data{
			"status": true,
			"diary": diary,
		})

		// old method and test
		//data, err := models.CreateDiary(db, diary)
		//
		//return c.JSON(http.StatusCreated, data)
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, data{
			"status": true,
		})
	}
}