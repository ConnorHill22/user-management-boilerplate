package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ConnorHill22/user-management-boilerplate/settings"
	sup "github.com/ConnorHill22/user-management-boilerplate/supabase"
	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"
)

// main function
func main() {
	// Initalize Supabase
	sup.Init(settings.SupabaseURL, settings.SupabaseKey)

	ctx := context.Background()
	user, err := sup.Client.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    "example@example.com",
		Password: "password",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	// create a new echo instance
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(settings.BaseURL))
}
