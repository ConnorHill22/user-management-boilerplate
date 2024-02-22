package view_home_route

import (
	"github.com/labstack/echo/v4"
)

func GetHomSubroutes(e *echo.Echo) {
	e.GET("/", HomeEndpoint)
}
