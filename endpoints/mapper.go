package endpoints

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type mapper struct {
	db *database.Queries
	e  *echo.Echo
	bc *bestchange.Client
}

func Create(e *echo.Echo, d *database.Queries, b *bestchange.Client) *mapper {
	m := &mapper{
		db: d,
		e:  e,
		bc: b,
	}
	e.Static("/", "dist")

	api := e.Group("/api")
	api.GET("/order", func(c echo.Context) error {
		return c.String(http.StatusOK, "users")
	})
	api.GET()

	return m
}

func (m *mapper) Run(port string) error {
	m.e.Logger.Fatal(m.e.Start("localhost:" + port))
	return errors.New("unable to run anymore")
}
