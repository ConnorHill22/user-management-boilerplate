package view_auth_route

import "github.com/labstack/echo/v4"

func GetAuthSubroutes(group *echo.Group) {
	group.GET("/login", SignInEndpoint)
}
