package router

import (
	"context"
	"lit-go/views"

	"github.com/labstack/echo/v4"
)

func InitRouter() {
	e := echo.New()
	e.Static("/dist", "./dist")
	e.Static("/styles", "./styles")
	api := e.Group("/api")

	e.GET("/", func(c echo.Context) error {
		return views.Index().Render(context.Background(), c.Response())
	})

	// API routes

	api.GET("/hello", getHello)

	e.Logger.Fatal(e.Start(":3000"))
}
