package api

import (
	"dreamers/api/handlers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func APIGroup(g *echo.Group, db *gorm.DB) {
	g.GET("diaries", handlers.GetDiaries(db))
	g.POST("diaries", handlers.CreateDiary(db))

	g.POST("accounts/login", handlers.Login(db))
}
