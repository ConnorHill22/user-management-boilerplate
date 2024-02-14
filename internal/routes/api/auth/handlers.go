package api_auth_route

import (
	"encoding/json"
	"log"
	"net/http"

	sup "github.com/ConnorHill22/user-management-boilerplate/internal/supabase"
	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"
)

type UserCredentials struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func RegisterUserEndpoint(c echo.Context) error {
	creds := UserCredentials{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	user, err := sup.Client.Auth.SignUp(c.Request().Context(), supa.UserCredentials{
		Email:    creds.Email,
		Password: creds.Password,
	})
	if err != nil {
		panic(err)
	}
	log.Println(user)
	return c.String(http.StatusOK, "User Registered")
}

func SignInEndpoint(c echo.Context) error {
	creds := UserCredentials{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	user, err := sup.Client.Auth.SignIn(c.Request().Context(), supa.UserCredentials{
		Email:    creds.Email,
		Password: creds.Password,
	})

	if err != nil {
		panic(err)
	}
	log.Println(user)
	return c.String(http.StatusOK, "User SignedIn")
}
