package api_auth_route

import "github.com/labstack/echo/v4"

func GetAuthSubroutes(group *echo.Group) {
	group.POST("/register", RegisterUserEndpoint)
	group.POST("/login", SignInEndpoint)
}
