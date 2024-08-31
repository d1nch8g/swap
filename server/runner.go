package server

import (
	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

// add endpoint to give exchangers info for bestchange

func Run(dir, port, tls string, e *echo.Echo, d *database.Queries, b *bestchange.Client) {
	e.Static("/", dir)

	api := e.Group("/api")

	orderservice := &orderservice{
		db: d,
		e:  e,
		bc: b,
	}
	api.POST("/neworder", orderservice.CreateOrder)
	api.GET("/params/:give/:receive", orderservice.ActualParams)

	api.GET("/neworder/*", orderservice.CreateOrder)
	api.GET("/info", orderservice.CreateOrder)

	if tls != "" {
		e.Logger.Fatal(e.StartAutoTLS(tls))
	}
	e.Logger.Fatal(e.Start("localhost:" + port))
}
