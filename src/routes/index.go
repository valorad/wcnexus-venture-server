package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var mainGroup *echo.Group

func UseRoute(group *echo.Group, routes func(group *echo.Group)) {
	routes(group)
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "api works!")
}

func IndexRoutes(group *echo.Group) {
	group.GET("/a", Index)
}

func ActivateIndex(g *echo.Group) {
	mainGroup = g
	UseRoute(mainGroup, IndexRoutes)

	// sub routes

}
