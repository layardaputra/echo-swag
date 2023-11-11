package user

import "github.com/labstack/echo/v4"

type UserResponse struct {
	Message string `json:"message"`
}

// GetUserHandler godoc
// @Summary      Get User Data
// @Description  get user data
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object} UserResponse
// @Router       /v1/users [get]
func GetUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		result := UserResponse{
			Message: "You got the user",
		}
		return c.JSON(200, result)
	}
}
