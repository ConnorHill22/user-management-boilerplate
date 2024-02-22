package view_dashboard_route

import (
	"github.com/ConnorHill22/user-management-boilerplate/internal/middleware"
	"github.com/labstack/echo/v4"
)

func GetDashboardSubroutes(e *echo.Echo) {
	e.GET("/dashboard", DashboardEndpoint, middleware.ValidateToken())
}
