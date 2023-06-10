package routers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouters() http.Handler {
	e := echo.New()

	// middleware section
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURI:       true,
		LogRemoteIP:  true,
		LogRequestID: true,
		LogMethod:    true,
		LogUserAgent: true,
		LogRoutePath: true,
		LogHost:      true,
		BeforeNextFunc: func(c echo.Context) {
			if c.Request().Header.Get(echo.HeaderXRequestID) == "" {
				c.Request().Header.Set(echo.HeaderXRequestID, uuid.NewString())
			}
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("[%v] REQUEST: uri: %v, Host: %v, Method: %v, UserAgent: %v, RoutePath: %v, IP: %v\n", v.RequestID, v.URI, v.Host, v.Method, v.UserAgent, v.RoutePath, v.RemoteIP)
			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Hello World!"})
	})

	return e
}
