package server

import (
	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type OrderService struct {
	db *database.Queries
	e  *echo.Echo
	bc *bestchange.Client
}

func (m *OrderService) ActualParams(c echo.Context) error {

	// min max exchange rate
	return nil
}

func (m *OrderService) CreateOrder(c echo.Context) error {
	return nil
}
