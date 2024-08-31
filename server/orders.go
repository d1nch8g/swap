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

// type CreateOrderRequest struct{
// 	User
// }

// This function is used to create order
func (m *OrderService) CreateOrder(c echo.Context) error {

	// m.bc.EstimateOperation()
	return nil
}
