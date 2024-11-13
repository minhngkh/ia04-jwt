package route

import (
	"test-echo/internal/auth"
	"test-echo/internal/misc"
	"test-echo/internal/user"

	"github.com/labstack/echo/v4"
)

func RegisterRootRoutes(e *echo.Group) {
	e.GET("/", misc.HelloWorldHandler)
	e.GET("/health", misc.DBHealthHandler)

	e.POST("/register", auth.RegisterHandler)
	e.POST("/login", auth.LoginHandler)

	e.GET("/profile", user.ProfileHandler, auth.JwtMiddleware())
}
