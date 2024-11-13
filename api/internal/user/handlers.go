package user

import (
	"net/http"
	"test-echo/internal/auth"

	"github.com/labstack/echo/v4"
)

func ProfileHandler(c echo.Context) error {
	user := auth.GetUserInfoFromJwtToken(c)

	info, err := GetUserInfo(user.Email)
	if err != nil {
		return echo.NewHTTPError(echo.ErrInternalServerError.Code, "Error getting user info")
	}

	return c.JSON(http.StatusOK, info)
}
