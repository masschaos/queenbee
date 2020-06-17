package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())

	return e
}
