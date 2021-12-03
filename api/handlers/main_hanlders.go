package handlers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/foolin/echo-template"
)

func Index(c echo.Context) error {
	return echotemplate.Render(c, http.StatusOK, "index", echo.Map{
		"title": "Garden of Stars",
	})
}