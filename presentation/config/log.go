package config

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
)

func SetupLog(e *echo.Echo) {
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("%s\n", reqBody)
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
