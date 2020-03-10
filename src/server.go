package main

import (
	"github.com/labstack/echo/v4"

	"github.com/valorad/wcnexus-venture-server/src/routes"
)

func main() {
	e := echo.New()

	apiGroup := e.Group("/api")

	routes.ActivateIndex(apiGroup)

	e.Logger.Fatal(e.Start(":1323"))
}
