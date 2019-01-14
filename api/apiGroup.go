package api

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"sample_echo/api/handlers"
)

func APIGroup(g *echo.Group, db *gorm.DB) {
	g.GET("diaries", handlers.GetDiaries(db))
	g.POST("diaries", handlers.CreateDiary(db))

	g.POST("accounts/login", handlers.Login(db))
}