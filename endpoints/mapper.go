package endpoints

import (
	"errors"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type mapper struct {
	*database.Queries
	*echo.Echo
}

func Create(e *echo.Echo, d *database.Queries) *mapper {
	e.Static("/", "public")

	m := &mapper{
		Queries: d,
		Echo:    e,
	}

	e.POST("/order/{from}/{to}", func(c echo.Context) error {
		return nil
	})

	return m
}

func (m *mapper) Run(port string) error {
	m.Echo.Logger.Fatal(m.Echo.Start(port))
	return errors.New("unable to run anymore")
}
