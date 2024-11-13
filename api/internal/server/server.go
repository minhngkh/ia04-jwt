package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"test-echo/internal/db"
	"test-echo/internal/logger"
	"test-echo/internal/route"
	"test-echo/internal/validator"
)

func NewEchoHandler() *echo.Echo {
	// db connection
	db.Get()

	e := echo.New()

	e.Validator = validator.New()

	e.Use(logger.LogWithZerolog())
	e.Use(echoMiddleware.Recover())

	rootGroup := e.Group("")
	route.RegisterRootRoutes(rootGroup)

	return e
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	e := NewEchoHandler()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      e,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
