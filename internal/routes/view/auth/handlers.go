package view_auth_route

import (
	"net/http"

	lib "github.com/ConnorHill22/user-management-boilerplate/library"
	web_auth "github.com/ConnorHill22/user-management-boilerplate/web/view/auth"
	"github.com/labstack/echo/v4"
)

func SignInEndpoint(c echo.Context) error {
	return lib.Render(c, http.StatusOK, web_auth.Login())
}
