package book

import "github.com/labstack/echo/v4"

type BookResponse struct {
	Message string `json:"message"`
}

// GetBookHandler godoc
// @Summary      Get Book Data
// @Description  get book data
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object} BookResponse
// @Router       /v1/books [get]
func GetBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		books := BookResponse{
			Message: "You got the book",
		}
		return c.JSON(200, books)
	}
}
