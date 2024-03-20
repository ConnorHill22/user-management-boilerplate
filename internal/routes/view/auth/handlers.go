package view_auth_route

import (
	"fmt"
	"log"
	"net/http"

	sup "github.com/ConnorHill22/user-management-boilerplate/internal/supabase"
	lib "github.com/ConnorHill22/user-management-boilerplate/library"
	"github.com/ConnorHill22/user-management-boilerplate/web/component"
	"github.com/a-h/templ"

	view_dashboard_route "github.com/ConnorHill22/user-management-boilerplate/internal/routes/view/dashboard"
	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"
)

type UserCredentials struct {
	Email    string `form:"Email"`
	Password string `form:"Password"`
}

type NewUserCredentials struct {
	Email           string `form:"Email"`
	Password        string `form:"Password"`
	ConfirmPassword string `form:"Confirm Password"`
}

type AuthLayoutTemplateProps struct {
	Title                string
	Subtitle             string
	SwitchAuthPrefixText string
	SwitchAuthPath       string
	SwitchAuthActionText string
	FormComponents       templ.Component
	Message              templ.Component
}

//
// Sign In Handler
//

var defaultSignInProps AuthLayoutTemplateProps = AuthLayoutTemplateProps{
	Title:                "Welcome Back!",
	Subtitle:             "Please login",
	SwitchAuthPrefixText: "Don't have an account? ",
	SwitchAuthPath:       "/auth/signup",
	SwitchAuthActionText: "Register now!",
	FormComponents:       SignInFormTemplate(),
	Message:              component.BlankMessage(),
}

func SignInGetHandler(c echo.Context) error {
	return lib.Render(c, http.StatusOK, AuthLayoutTemplate(defaultSignInProps))
}

func SignInPostHandler(c echo.Context) error {
	creds := &UserCredentials{}
	props := defaultSignInProps

	if err := c.Bind(creds); err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		props.Message = component.SimpleError("Server Error. Please try again")
		return lib.Render(c, http.StatusInternalServerError, AuthLayoutTemplate(props))
	}
	user, err := sup.Client.Auth.SignIn(c.Request().Context(), supa.UserCredentials{
		Email:    creds.Email,
		Password: creds.Password,
	})
	if err != nil {
		props.Message = component.SimpleError("Invalid Credentials")
		return lib.Render(c, http.StatusInternalServerError, AuthLayoutTemplate(props))
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = user.AccessToken
	cookie.Path = "/"
	c.SetCookie(cookie)

	return view_dashboard_route.DashboardHandler(c)
}

//
// Sign Up Handler
//

var defaultSignUpProps AuthLayoutTemplateProps = AuthLayoutTemplateProps{
	Title:                "Welcome to AIdaptive",
	Subtitle:             "Please register below",
	SwitchAuthPrefixText: "Already have an account? ",
	SwitchAuthPath:       "/auth/signin",
	SwitchAuthActionText: "Click to Login.",
	FormComponents:       SignUpFormTemplate(),
	Message:              component.BlankMessage(),
}

func SignUpGetHandler(c echo.Context) error {
	return lib.Render(c, http.StatusOK, AuthLayoutTemplate(defaultSignUpProps))
}

func SignUpPostHandler(c echo.Context) error {
	creds := &NewUserCredentials{}
	props := defaultSignUpProps
	if err := c.Bind(creds); err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		props.Message = component.SimpleError("Server Error. Please try again")
		return lib.Render(c, http.StatusInternalServerError, AuthLayoutTemplate(props))
	}

	if creds.Password != creds.ConfirmPassword {
		props.Message = component.SimpleError("Passwords don't match")
		return lib.Render(c, http.StatusInternalServerError, AuthLayoutTemplate(props))
	}

	user, err := sup.Client.Auth.SignUp(c.Request().Context(), supa.UserCredentials{
		Email:    creds.Email,
		Password: creds.Password,
	})

	if err != nil {
		props.Message = component.SimpleError(err.Error())
		return lib.Render(c, http.StatusInternalServerError, AuthLayoutTemplate(props))
	}

	props.Message = component.SimpleMessage(fmt.Sprintf(
		"User has been created. Please check your email (%s) and click the verification link",
		user.Email,
	))
	return lib.Render(c, http.StatusInternalServerError, AuthLayoutTemplate(props))
}
