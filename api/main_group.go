package api

import (
	"dreamers/api/handlers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func MainGroup(g *echo.Group, db *gorm.DB) {
	g.GET("", handlers.Index)
	//g.GET("/diaries", handlers.GetDiaries(db))
	//g.POST("/diaries", handlers.CreateDiary(db))

	//e.GET("/", handlers.Index)
	//e.GET("/diaries", handlers.GetDiaries(db))
	//e.POST("/diaries", handlers.CreateDiary(db))
}
