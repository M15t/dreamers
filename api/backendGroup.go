package api

import (
	"dreamers/api/handlers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func BackendGroup(g *echo.Group, db *gorm.DB) {
	g.GET("/", handlers.Manage)
}
