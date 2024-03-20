package view_dashboard_route

import (
	// "errors"
	"net/http"

	lib "github.com/ConnorHill22/user-management-boilerplate/library"
	"github.com/a-h/templ"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ViewName string

const (
	Home        ViewName = "home"
	Tickets     ViewName = "tickets"
	InstantHelp ViewName = "instantHelp"
	Metrics     ViewName = "metrics"
	Settings    ViewName = "settings"
)

type DashboardIcon struct {
	Name ViewName
	Icon string
	Path string
	Alt  string
}

var MainDashboardIcons = [4]DashboardIcon{
	{Name: Home, Icon: "/static/icons/home_icon.svg", Path: "/dashboard", Alt: "Home"},
	{Name: Tickets, Icon: "/static/icons/ticket_icon.svg", Path: "/dashboard/tickets", Alt: "Tickets"},
	{Name: InstantHelp, Icon: "/static/icons/fire_icon.svg", Path: "/dashboard/instantHelp", Alt: "Instant Help"},
	{Name: Metrics, Icon: "/static/icons/metrics_icon.svg", Path: "/dashboard/metrics", Alt: "Metrics"},
}

type DashboardData struct {
	ActiveView     ViewName
	DashboardIcons []DashboardIcon
	UserId         string
	Component      templ.Component
}

func DashboardHandler(c echo.Context) error {
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["sub"].(string)

	dashboardData := DashboardData{
		ActiveView:     Home,
		DashboardIcons: MainDashboardIcons[:],
		UserId:         "1234",
		Component:      HomeTemplate(),
	}
	return lib.Render(c, http.StatusOK, DashboardTemplate(dashboardData))
}

func TicketsHandler(c echo.Context) error {
	dashboardData := DashboardData{
		ActiveView:     Tickets,
		DashboardIcons: MainDashboardIcons[:],
		UserId:         "1234",
		Component:      HomeTemplate(),
	}
	return lib.Render(c, http.StatusOK, DashboardTemplate(dashboardData))
}

func UserSettingsHandler(c echo.Context) error {
	dashboardData := DashboardData{
		ActiveView:     Settings,
		DashboardIcons: MainDashboardIcons[:],
		UserId:         "1234",
		Component:      UserSettingsTemplate(),
	}
	return lib.Render(c, http.StatusOK, DashboardTemplate(dashboardData))
}

// func InstantHelpHandler(c echo.Context) error {
// 	return lib.Render(c, http.StatusOK, web_dashboard.InstantHelp())
// }

// func MetricsHandler(c echo.Context) error {
// 	return lib.Render(c, http.StatusOK, web_dashboard.Metrics())
// }
