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

	userSvc := &UserService{
		db: d,
		e:  e,
		bc: b,
	}

	api.POST("/login", userSvc.Login)
	

	if tls != "" {
		e.Logger.Fatal(e.StartAutoTLS(tls))
	}
	e.Logger.Fatal(e.Start("localhost:" + port))
}
