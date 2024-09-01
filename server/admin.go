package server

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

// @Summary	Login to platform user account
// @ID			login
// @Accept		json
// @Produce	json
// @Param		Email		header		string	true	"Email login"
// @Param		Password	header		string	true	"Password"
// @Success	200			{string}	string	"ok"
// @Failure	401			{object}	string	"Unautharized"
// @Router		/login [post]
func (s *Endpoints) Login(c echo.Context) error {
	email := c.Request().Header["Email"]
	password := c.Request().Header["Password"]

	if email == nil || password == nil {
		c.Response().WriteHeader(http.StatusUnauthorized)
		_, err := c.Response().Write([]byte("empty login or password"))
		return err
	}

	user, err := s.db.GetUser(c.Request().Context(), email[0])
	if err != nil {
		c.Response().WriteHeader(http.StatusUnauthorized)
		return errors.New("unable to login")
	}

	hasher := sha512.New()
	hasher.Write([]byte(password[0]))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if user.Passwhash != sha {
		c.Response().WriteHeader(http.StatusUnauthorized)
		_, err := c.Response().Write([]byte("bad password"))
		return err
	}

	tokenhasher := sha512.New()
	tokenhasher.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	token := base64.URLEncoding.EncodeToString(tokenhasher.Sum(nil))

	_, err = s.db.UpdateUserToken(c.Request().Context(), database.UpdateUserTokenParams{
		ID:    user.ID,
		Token: token,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusUnauthorized)
		_, err := c.Response().Write([]byte("unable to update token"))
		return err
	}

	_, err = c.Response().Write([]byte(token))
	return err
}

// verify email address

type Orders struct {
	ActiveOrders []database.Order `json:"orders"`
}

// @Summary	Get active orders as administrator accout
// @ID			admin.getorders
// @Accept		json
// @Produce	json
// @Success	200	{object}	Orders	"Orders"
// @Security	ApiKeyAuth
// @Router		/admin/getorders [get]
func (s *Endpoints) GetOrders(c echo.Context) error {
	orders, err := s.db.OrdersUnfinished(c.Request().Context())
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &Orders{
		ActiveOrders: orders,
	})
}

// @Summary	Create new currency for exchangers
// @ID			admin.createcurrency
// @Accept		json
// @Produce	json
// @Param		Body	body		database.CreateCurrencyParams	true	"Create user request"
// @Success	200		{object}	Orders							"Orders"
// @Security	ApiKeyAuth
// @Router		/admin/createcurrency [post]
func (s *Endpoints) CreateCurrency(c echo.Context) error {
	var curr database.CreateCurrencyParams
	c.Bind(&curr)
	_, err := s.db.CreateCurrency(c.Request().Context(), curr)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}
	return nil
}

func (s *Endpoints) RemoveCurrency(c echo.Context) error {
	return nil
}

// create, remove and lists currencies
// create, remove and list exchangers
// create and update balance
// execute order
