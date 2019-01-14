package api

import (
	"sample_echo/api/handlers"
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
)

func MainGroup(g *echo.Group, db *gorm.DB) {
	g.GET("", handlers.Index)
	//g.GET("/diaries", handlers.GetDiaries(db))
	//g.POST("/diaries", handlers.CreateDiary(db))

	//e.GET("/", handlers.Index)
	//e.GET("/diaries", handlers.GetDiaries(db))
	//e.POST("/diaries", handlers.CreateDiary(db))
}
