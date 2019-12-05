package controllers

import (
	"ddd-go/application"
	"ddd-go/domain"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func RouteFlight(e *echo.Echo) {
	e.POST("/v1/search/:vendorCode", func(c echo.Context) error {
		request := new(domain.SearchRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		vendorCode := c.Param("vendorCode", )
		//fmt.Println(time.Now(), "Start")
		status, response, errorResponse := application.Search(vendorCode, *request)
		if (status != http.StatusOK) {
			return c.JSON(status, errorResponse)
		}
		fmt.Println(time.Now(), "Finish")
		//time.Sleep(10 * time.Second)
		return c.JSON(status, response)
	})

	e.POST("/v1/bookInfo", func(c echo.Context) error {
		request := new(domain.BookInfoRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		status, response, errorResponse := application.BookInfo(*request)
		if (status != http.StatusOK) {
			return c.JSON(status, errorResponse)
		}
		return c.JSON(status, response)
	})
}
