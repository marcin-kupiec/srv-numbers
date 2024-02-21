package http

import "github.com/labstack/echo/v4"

func SetRoutes(e *echo.Echo, handler echo.HandlerFunc) {
	e.GET("/endpoint/:number", handler)
}
