package routers

import (
	"dreamers/api"
	"dreamers/database"
	"fmt"
	"net/http"

	"github.com/foolin/echo-template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/qor/validations"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.Render(http.StatusOK, errorPage, echo.Map{}); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func New() *echo.Echo {
	e := echo.New()

	// add default render views by echotemplate
	e.Renderer = echotemplate.Default()

	// add custom handle for 404, 500 error
	e.HTTPErrorHandler = customHTTPErrorHandler

	// init database
	db := database.InitDB("storage.db")

	// init validation
	validations.RegisterCallbacks(db)

	// define static
	e.Static("/static", "assets")

	// enable log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${method} ${status} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	// generate template for frontend
	mainMW := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:         "views/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		DisableCache: true,
	})

	// generate template for backend
	backendMW := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:         "views/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		DisableCache: true,
	})

	mainGroup := e.Group("/", mainMW)
	backendGroup := e.Group("/manage/", backendMW)
	apiGroup := e.Group("/api/")

	//middlewares.SetBackendMiddlewares(backendGroup)

	api.MainGroup(mainGroup, db)
	api.BackendGroup(backendGroup, db)
	api.APIGroup(apiGroup, db)

	return e
}
