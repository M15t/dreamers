package handlers

import (
	"github.com/labstack/echo"
	"github.com/foolin/echo-template"
	"net/http"
)

func Manage(c echo.Context) error {
	return echotemplate.Render(c, http.StatusOK, "index", echo.Map{
		"title": "Garden of Stars",
	})
}
