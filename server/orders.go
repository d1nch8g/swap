package server

import (
	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type orderservice struct {
	db *database.Queries
	e  *echo.Echo
	bc *bestchange.Client
}

func (m *orderservice) ActualParams(c echo.Context) error {

	// min max exchange rate
	return nil
}

func (m *orderservice) CreateOrder(c echo.Context) error {
	return nil
}
