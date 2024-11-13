package auth

import (
	"net/http"
	"test-echo/internal/validator"

	"github.com/labstack/echo/v4"
)

type registerRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}

type loginRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

func RegisterHandler(c echo.Context) error {
	req := new(registerRequest)
	if httpErr := validator.BindAndValidateRequest(c, req); httpErr != nil {
		return httpErr
	}

	if err := CreateUser(req.Email, req.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating user")
	}

	setJwtTokenInCookie(c, UserInfo{Email: req.Email})

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User successfully registered",
	})
}

func LoginHandler(c echo.Context) error {
	var req loginRequest
	if httpErr := validator.BindAndValidateRequest(c, &req); httpErr != nil {
		return httpErr
	}

	loginInfo := &LoginInfo{
		Email:    req.Email,
		Password: req.Password,
	}
	err := VerifyLoginInfo(loginInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	setJwtTokenInCookie(c, UserInfo{Email: req.Email})

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
	})
}

func setJwtTokenInCookie(c echo.Context, info UserInfo) error {
	token, err := NewJwtToken(info)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.SetCookie(&http.Cookie{
		Name:    "jwt",
		Value:   token.Value,
		Expires: token.ExpiresAt,
	})

	return nil
}
