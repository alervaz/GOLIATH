package router

import (
	"context"
	"lit-go/views"

	"github.com/labstack/echo/v4"
)

func getHello(c echo.Context) error {
	return views.Hello().Render(context.Background(), c.Response())
}
