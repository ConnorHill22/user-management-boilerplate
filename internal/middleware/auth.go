package middleware

import (
	"fmt"
	"net/http"

	"github.com/ConnorHill22/user-management-boilerplate/settings"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ValidateToken() echo.MiddlewareFunc {
	// JWT middleware configuration
	config := echojwt.Config{
		SigningKey:     []byte(settings.SupbabaseJwtSecret), // Supabase JWT secret
		SigningMethod:  "HS256",
		SuccessHandler: successHandler,
		ErrorHandler:   errorHandler,
		TokenLookup:    "cookie:token",
	}
	return echojwt.WithConfig(config)
}

// Success handler function for JWT middleware
func successHandler(c echo.Context) {
	user := c.Get("user")
	fmt.Println("User:", user)
}

// Error handler function for JWT middleware
func errorHandler(c echo.Context, err error) error {
	fmt.Println("Error:", err.Error())
	return c.String(http.StatusUnauthorized, err.Error())
}
