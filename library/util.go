package lib

import (
	"os"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Getenv(var_name string, default_value string) string {
	value := os.Getenv(var_name)
	if value == "" {
		value = default_value
	}
	return value
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
