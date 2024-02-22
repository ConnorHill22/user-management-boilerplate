package view_dashboard_route

import (
	// "errors"
	"fmt"
	"net/http"

	// lib "github.com/ConnorHill22/user-management-boilerplate/library"
	// web_dashboard "github.com/ConnorHill22/user-management-boilerplate/web/view/dashboard"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func DashboardEndpoint(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["sub"].(string)

	return c.String(http.StatusOK, fmt.Sprintf("Welcome, %s!", name))
	// return lib.Render(c, http.StatusOK, web_dashboard.Dashboard())
}
