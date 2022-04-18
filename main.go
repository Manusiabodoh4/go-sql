package main

import (
	routes "github.com/Manusiabodoh4/go-sql/src/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	routes.NewAccountRoutes(app.Group("/v1/account"))

	app.Logger.Fatal(app.Start(":4567"))

}
