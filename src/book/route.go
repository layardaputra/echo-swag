package book

import "github.com/labstack/echo/v4"

func RegisterBookRoute(e *echo.Echo) {
	e.GET("/v1/books/", GetBookHandler())
}
