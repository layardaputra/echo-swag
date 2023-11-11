package user

import "github.com/labstack/echo/v4"

func RegisterUserRoute(e *echo.Echo) {
	e.GET("/v1/users/", GetUserHandler())
}
