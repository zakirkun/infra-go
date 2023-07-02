package routers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	auth "github.com/zakirkun/infra-go/pkg/auth/jwt"
	"github.com/zakirkun/infra-go/pkg/config"
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
		return c.JSON(http.StatusOK, map[string]string{"messages": "Hello World!", "request-id": c.Request().Header.Get(echo.HeaderXRequestID)})
	})

	e.POST("/generate", func(c echo.Context) error {
		user := c.FormValue("user")
		jwtImpl := auth.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))

		form := map[string]interface{}{
			"user": user,
		}
		token, _ := jwtImpl.GenerateToken(form)

		response := map[string]string{
			"token":      token,
			"request-id": c.Request().Header.Get(echo.HeaderXRequestID),
		}

		return c.JSON(http.StatusOK, response)
	})

	e.POST("/validate", func(c echo.Context) error {
		token := c.FormValue("token")
		jwtImpl := auth.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
		valid, err := jwtImpl.ValidateToken(token)

		response := map[string]string{
			"request-id": c.Request().Header.Get(echo.HeaderXRequestID),
		}

		if err != nil {
			response["token"] = err.Error()
			return c.JSON(http.StatusBadRequest, response)
		}

		if !valid {
			response["token"] = fmt.Sprintf("%v token not valid >> %v", valid, token)
			return c.JSON(http.StatusUnauthorized, response)
		}

		response["token"] = fmt.Sprintf("%v token valid >> %v", valid, token)
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/parse", func(c echo.Context) error {
		token := c.FormValue("token")
		jwtImpl := auth.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
		valid, err := jwtImpl.ParseToken(token)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, valid)
	})

	return e
}
