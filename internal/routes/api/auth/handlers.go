package api_auth_route

import (
	"log"
	"net/http"

	sup "github.com/ConnorHill22/user-management-boilerplate/internal/supabase"
	lib "github.com/ConnorHill22/user-management-boilerplate/library"
	"github.com/ConnorHill22/user-management-boilerplate/web/view/component"
	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"
)

type UserCredentials struct {
	Email    string `form:"Email"`
	Password string `form:"Password"`
}

func RegisterUserEndpoint(c echo.Context) error {
	creds := new(UserCredentials)

	if err := c.Bind(creds); err != nil {
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
	creds := &UserCredentials{}

	if err := c.Bind(creds); err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	user, err := sup.Client.Auth.SignIn(c.Request().Context(), supa.UserCredentials{
		Email:    creds.Email,
		Password: creds.Password,
	})
	if err != nil {
		return lib.Render(c, http.StatusOK, component.SimpleError("Invalid Credentials"))
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = user.AccessToken
	cookie.Path = "/"
	c.SetCookie(cookie)
	c.Response().Header().Set("HX-Location", "/dashboard")
	return c.String(http.StatusOK, "User SignedIn")
}
