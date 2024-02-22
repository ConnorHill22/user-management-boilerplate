package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func DevLoggingMiddleware() echo.MiddlewareFunc {
	logger, _ := zap.NewDevelopment()
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("Method", v.Method),
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	})
}
