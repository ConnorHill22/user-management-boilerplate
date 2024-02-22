package main

import (
	"github.com/ConnorHill22/user-management-boilerplate/internal/middleware"
	"github.com/ConnorHill22/user-management-boilerplate/internal/routes"
	sup "github.com/ConnorHill22/user-management-boilerplate/internal/supabase"
	"github.com/ConnorHill22/user-management-boilerplate/settings"
	"github.com/labstack/echo/v4"
)

// main function
func main() {
	// Initalize Supabase
	sup.Init(settings.SupabaseURL, settings.SupabaseKey)

	e := echo.New()

	e.Use(middleware.DevLoggingMiddleware())

	e.Static("/static", "web/static")

	// Register Routes
	routes.RegisterViewRoutes(e)
	routes.RegisterApiRoutes(e.Group(("/api")))

	e.Logger.Fatal(e.Start(settings.BaseURL))
}
