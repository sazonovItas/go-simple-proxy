package router

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	prettylogger "github.com/rdbell/echo-pretty-logger"
)

func New(timeout time.Duration, usePrettyLogger bool) *echo.Echo {
	e := echo.New()

	e.Pre(middleware.AddTrailingSlash())

	e.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: timeout}))
	e.Use(middleware.RequestID())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if usePrettyLogger {
				return prettylogger.Logger(next)(c)
			}

			return middleware.Logger()(next)(c)
		}
	})

	return e
}
