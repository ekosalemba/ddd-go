package main

import (
	"ddd-go/presentation/config"
	"ddd-go/presentation/controllers"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	cfg, _ := config.LoadConfiguration()
	fmt.Println("Starting application...")
	e := echo.New()
	config.SetupLog(e)
	controllers.RouteFlight(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/api-docs", "public/swaggerui")
	e.Static("/api-spec", "swagger.yml")
	e.Logger.Fatal(e.Start(":" + cfg.Application.Port))
}
