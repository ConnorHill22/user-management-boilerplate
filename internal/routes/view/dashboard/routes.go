package view_dashboard_route

import "github.com/labstack/echo/v4"

func GetDashboardSubroutes(e *echo.Echo) {
	e.GET("/dashboard", DashboardHandler)
	e.GET("/dashboard/tickets", TicketsHandler)
	// e.GET("/dashboard/instantHelp", InstantHelpHandler)
	// e.GET("/dashboard/metrics", MetricsHandler)
	e.GET("/dashboard/user/:userId/settings", UserSettingsHandler)

	// e.GET("/dashboard", DashboardHandler, middleware.ValidateToken())
	// e.GET("/dashboard/tickets", DashboardHandler, middleware.ValidateToken())
	// e.GET("/dashboard/instantHelp", DashboardHandler, middleware.ValidateToken())
	// e.GET("/dashboard/metrics", DashboardHandler, middleware.ValidateToken())
}
