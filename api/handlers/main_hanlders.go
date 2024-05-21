package handlers

import (
	"net/http"

	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
)

// Index is a function that handles the rendering of the "index" template with a title of "Garden of Stars" using echoview package for Echo framework.
func Index(c echo.Context) error {
	return echoview.Render(c, http.StatusOK, "index", echo.Map{
		"title": "Garden of Stars",
	})
}
