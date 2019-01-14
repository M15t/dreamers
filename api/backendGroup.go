package api

import (
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"sample_echo/api/handlers"
)

func BackendGroup(g *echo.Group, db *gorm.DB) {
	g.GET("/", handlers.Manage)
}
