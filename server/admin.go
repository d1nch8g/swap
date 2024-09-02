package server

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/inswap/gen/database"
)

func (e *Endpoints) Login(c echo.Context) error {
	email := c.Request().Header["Email"]
	password := c.Request().Header["Password"]

	if email == nil || password == nil {
		c.Response().WriteHeader(http.StatusUnauthorized)
		_, err := c.Response().Write([]byte("empty login or password"))
		return err
	}

	user, err := e.db.GetUser(c.Request().Context(), email[0])
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

	_, err = e.db.UpdateUserToken(c.Request().Context(), database.UpdateUserTokenParams{
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

type Busy struct {
	Busy bool `json:""`
}

func (e *Endpoints) ChangeBusy(c echo.Context) error {
	a := c.Request().Header["Token"]

	u, err := e.db.GetUserByToken(c.Request().Context(), a[0])
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	u, err = e.db.UpdateUserBusy(c.Request().Context(), database.UpdateUserBusyParams{
		Email: u.Email,
		Busy:  !u.Busy,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &Busy{
		Busy: u.Busy,
	})
}

type Orders struct {
	ActiveOrders []database.Order `json:"orders"`
}

func (e *Endpoints) GetOrders(c echo.Context) error {
	orders, err := e.db.OrdersUnfinished(c.Request().Context())
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &Orders{
		ActiveOrders: orders,
	})
}

func (e *Endpoints) CreateCurrency(c echo.Context) error {
	var curr database.CreateCurrencyParams
	err := c.Bind(&curr)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to get unmarshal request"))
		return err
	}

	_, err = e.db.CreateCurrency(c.Request().Context(), curr)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}
	return nil
}

type RemoveCurrencyRequest struct {
	Code string `json:"code"`
}

func (e *Endpoints) RemoveCurrency(c echo.Context) error {
	var curr RemoveCurrencyRequest
	err := c.Bind(&curr)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to get unmarshal request"))
		return err
	}

	err = e.db.RemoveCurrency(c.Request().Context(), curr.Code)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

func (e *Endpoints) CreateExchanger(c echo.Context) error {
	var createExchanger database.CreateExchangerParams
	err := c.Bind(&createExchanger)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to get unmarshal request"))
		return err
	}

	_, err = e.db.CreateExchanger(c.Request().Context(), createExchanger)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

// create, remove and list exchangers
// create and update balance
// execute order
