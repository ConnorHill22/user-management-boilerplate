package view_auth_route

import "github.com/labstack/echo/v4"

func GetAuthSubroutes(group *echo.Group) {

	group.GET("/signin", SignInGetHandler)
	group.POST("/signin", SignInPostHandler)

	group.GET("/signup", SignUpGetHandler)
	group.POST("/signup", SignUpPostHandler)
}
