package view_home_route

import (
	"net/http"

	lib "github.com/ConnorHill22/user-management-boilerplate/library"
	web_home "github.com/ConnorHill22/user-management-boilerplate/web/view/home"
	"github.com/labstack/echo/v4"
)

func HomeEndpoint(c echo.Context) error {
	return lib.Render(c, http.StatusOK, web_home.Home())
}
