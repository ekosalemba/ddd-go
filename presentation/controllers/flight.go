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
		if status != http.StatusOK {
			return c.JSON(status, errorResponse)
		}
		//go asyncExample()
		//time.Sleep(5 * time.Second)
		fmt.Println(time.Now(), "Finish Response")
		return c.JSON(status, response)
	})

	e.POST("/v1/book", func(c echo.Context) error {
		request := new(domain.BookRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		status, response, errorResponse := application.Book(*request)
		if status != http.StatusOK {
			return c.JSON(status, errorResponse)
		}
		return c.JSON(status, response)
	})

	e.POST("/v1/bookInfo", func(c echo.Context) error {
		request := new(domain.BookInfoRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		status, response, errorResponse := application.BookInfo(*request)
		if status != http.StatusOK {
			return c.JSON(status, errorResponse)
		}
		return c.JSON(status, response)
	})

	e.POST("/v1/setPayment", func(c echo.Context) error {
		request := new(domain.SetPaymentRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		status, response, errorResponse := application.SetPayment(*request)
		if status != http.StatusOK {
			return c.JSON(status, errorResponse)
		}
		return c.JSON(status, response)
	})
}
func asyncExample() {

	time.Sleep(10 * time.Second)
	fmt.Println(time.Now(), "Finish Async Process")

}
