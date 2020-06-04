package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/masschaos/queenbee/apis"
)

func router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())

	template := e.Group("templates")
	{
		template.POST("", apis.CreateTemplate)
		template.GET("/:id", apis.ReadTemplate)
		template.GET("", apis.ListTemplate)
		template.PATCH("/:id", apis.UpdateTemplate)
		template.DELETE("/:id", apis.DeleteTemplate)
	}

	job := e.Group("jobs")
	{
		job.POST("", apis.CreateJob)
		job.GET("", apis.ReadJob)
		job.POST("/:id/result", apis.CreateJobResult)
	}

	return e
}
