package handlers

import (
	"net/http"

	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
)

// Manage renders the "index" template with a title of "Garden of Stars" using echoview package and returns an error if any.
func Manage(c echo.Context) error {
	return echoview.Render(c, http.StatusOK, "index", echo.Map{
		"title": "Garden of Stars",
	})
}
