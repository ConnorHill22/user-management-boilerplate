package main

import (
	"github.com/ConnorHill22/user-management-boilerplate/internal/routes"
	sup "github.com/ConnorHill22/user-management-boilerplate/internal/supabase"
	"github.com/ConnorHill22/user-management-boilerplate/settings"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// main function
func main() {
	// Initalize Supabase
	sup.Init(settings.SupabaseURL, settings.SupabaseKey)

	e := echo.New()

	logger, _ := zap.NewProduction()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))

	e.Static("/static", "web/static")

	// Register Routes
	routes.RegisterViewRoutes(e)
	routes.RegisterApiRoutes(e.Group(("/api")))

	e.Logger.Fatal(e.Start(settings.BaseURL))
}
