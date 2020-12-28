package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hasura/go-graphql-client"
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

func GetTheme(c echo.Context) error {

	client := graphql.NewClient("https://example.com/graphql", nil)

	var query struct {
		Api struct {
			Name graphql.String
		}
	}

	err := client.Query(context.Background(), &query, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c.String(http.StatusOK, string(query.Api.Name))

}

func ThemeRoutes(group *echo.Group) {
	group.GET("/thm", GetTheme)
}

func ActivateIndex(g *echo.Group) {
	mainGroup = g
	UseRoute(mainGroup, IndexRoutes)
	UseRoute(mainGroup, ThemeRoutes)

	// sub routes

}
