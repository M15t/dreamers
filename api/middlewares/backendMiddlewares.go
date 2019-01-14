package middlewares

import (
	"github.com/labstack/echo"
)

func SetBackendMiddlewares(g *echo.Group) {
	//g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	//	if username == "joe" && password == "secret" {
	//		return true, nil
	//	}
	//	return false, nil
	//}))
}
