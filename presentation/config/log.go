package config

import (
	"bytes"
	"ddd-go/presentation/utils"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func SetupLog(e *echo.Echo) {
	e.Use(middlewareRequest)
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		saveLogResponse(c, resBody)
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func middlewareRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		saveLogRequest(c)
		return next(c)
	}
}
func saveLogRequest(c echo.Context) {
	logType := "request"
	logFileName := logType + "_" + time.Now().Format("2006-01-02") + ".log"
	logFile, err := os.OpenFile("log/"+logType+"/"+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
			reqGuid := strings.TrimSpace(string(utils.GenerateGuid()))
			c.Request().Header.Set("request-guid", reqGuid)
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			logger := log.New(logFile, "", log.LstdFlags)
			currentTimestamp := time.Now().Format("2006-01-02 15:04:05")
			logger.Println(fmt.Sprintf("{'guid': '%s', 'timestamp': '%s', 'method': '%s', 'path': '%s', 'body': %s}",
				reqGuid, currentTimestamp, c.Request().Method, c.Request().RequestURI, string(bodyBytes)))
		}
	}
}
func saveLogResponse(c echo.Context, responseBody []byte) {
	logType := "response"
	logFileName := logType + "_" + time.Now().Format("2006-01-02") + ".log"
	logFile, err := os.OpenFile("log/"+logType+"/"+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		logger := log.New(logFile, "", log.LstdFlags)
		reqGuid := c.Request().Header.Get("request-guid")
		currentTimestamp := time.Now().Format("2006-01-02 15:04:05")
		logger.Println(fmt.Sprintf("{'guid': '%s', 'timestamp': '%s', 'method': '%s', 'path': '%s', 'body': %s}",
			reqGuid, currentTimestamp, c.Request().Method, c.Request().RequestURI, responseBody))
	}
}
