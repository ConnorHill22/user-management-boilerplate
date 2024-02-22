package routes

import (
	api_auth_route "github.com/ConnorHill22/user-management-boilerplate/internal/routes/api/auth"
	view_auth_route "github.com/ConnorHill22/user-management-boilerplate/internal/routes/view/auth"
	view_dashboard_route "github.com/ConnorHill22/user-management-boilerplate/internal/routes/view/dashboard"
	view_home_route "github.com/ConnorHill22/user-management-boilerplate/internal/routes/view/home"
	"github.com/labstack/echo/v4"
)

func RegisterApiRoutes(group *echo.Group) {
	api_auth_route.GetAuthSubroutes(group.Group(("/auth")))
}

func RegisterViewRoutes(e *echo.Echo) {
	view_home_route.GetHomSubroutes(e)
	view_dashboard_route.GetDashboardSubroutes(e)
	view_auth_route.GetAuthSubroutes((e.Group("/auth")))
}
