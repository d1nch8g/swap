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

type CreateOrderRequest struct {
	Email  string  `json:"email"`
	Input  string  `json:"input"`
	Ouput  string  `json:"output"`
	Amount float64 `json:"amount"`
}

//	@Summary	Create new order
//	@ID			order.create
//	@Accept		json
//	@Produce	json
//	@Param		Body	body		CreateOrderRequest	true "Create order body"
//	@Success	200		{string}	string				"ok"
//	@Router		/createorder [post]
func (m *OrderService) CreateOrder(c echo.Context) error {
	us := &UserService{
		db: m.db,
		e:  m.e,
		bc: m.bc,
	}
	err := us.CreateUser(c)
	if err != nil {
		return err
	}

	// Estimate exchange rate for order and create it for user with one free admin
	return nil
}
