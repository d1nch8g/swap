package server

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type CreateUserRequest struct {
	Email string `json:"email"`
}

// @Summary	Create new user
// @ID			user.create
// @Accept		json
// @Produce	json
// @Param		Body	body		CreateUserRequest	true	"Create user request"
// @Success	200		{object}	Orders				"ok"
// @Router		/createuser [get]
func (s *Endpoints) CreateUser(c echo.Context) error {
	var createUser CreateUserRequest
	err := c.Bind(&createUser)
	if err != nil {
		c.Response().WriteHeader(http.StatusForbidden)
		_, err := c.Response().Write([]byte("unable to bind user request"))
		return err
	}

	_, err = s.db.CreateUser(c.Request().Context(), database.CreateUserParams{
		Email:     createUser.Email,
		Passwhash: "nil",
		Token:     "nil",
	})
	if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint ") {
		c.Response().WriteHeader(http.StatusForbidden)
		_, err := c.Response().Write([]byte("unable to create new user"))
		return err
	}

	return nil
}

type CreateOrderRequest struct {
	Email  string  `json:"email"`
	Input  string  `json:"input"`
	Ouput  string  `json:"output"`
	Amount float64 `json:"amount"`
}

// @Summary	Create new order
// @ID			order.create
// @Accept		json
// @Produce	json
// @Param		Body	body		CreateOrderRequest	true	"Create order body"
// @Success	200		{string}	string				"ok"
// @Router		/createorder [post]
func (m *Endpoints) CreateOrder(c echo.Context) error {
	err := m.CreateUser(c)
	if err != nil {
		return err
	}

	// Estimate exchange rate for order and create it for user with one free admin
	return nil
}

// verify card
// payment approve
