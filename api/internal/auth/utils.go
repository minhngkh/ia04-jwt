package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func GetUserInfoFromJwtToken(c echo.Context) *UserInfo {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtClaims)
	log.Info().Msg(claims.Email)

	return &claims.UserInfo
}
