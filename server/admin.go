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
	Busy bool `json:"busy"`
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
	token := c.Request().Header["Token"][0]

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	orders, err := e.db.GetOrders(c.Request().Context(), u.ID)
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
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
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
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
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
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
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
func (e *Endpoints) RemoveExchanger(c echo.Context) error {
	var rmexchanger database.RemoveExchangerParams
	err := c.Bind(&rmexchanger)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	err = e.db.RemoveExchanger(c.Request().Context(), rmexchanger)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

type CreateBalanceRequest struct {
	CurrencyId int64
	Balance    float64
	Address    string
}

func (e *Endpoints) CreateBalance(c echo.Context) error {
	var req CreateBalanceRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := c.Request().Header["Token"][0]

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	_, err = e.db.CreateBalance(c.Request().Context(), database.CreateBalanceParams{
		UserID:     u.ID,
		CurrencyID: req.CurrencyId,
		Balance:    req.Balance,
		Address:    req.Address,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

type UpdateBalanceRequest struct {
	BalanceId int64   `json:"balance_id"`
	Balance   float64 `json:"balance"`
}

func (e *Endpoints) UpdateBalance(c echo.Context) error {
	var req UpdateBalanceRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := c.Request().Header["Token"][0]

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	_, err = e.db.UpdateBalance(c.Request().Context(), database.UpdateBalanceParams{
		ID:      req.BalanceId,
		UserID:  u.ID,
		Balance: req.Balance,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}
	return nil
}

type Balances struct {
	Balances []database.Balance
}

func (e *Endpoints) ListBalances(c echo.Context) error {
	token := c.Request().Header["Token"][0]

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	balances, err := e.db.ListBalances(c.Request().Context(), u.ID)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &Balances{
		Balances: balances,
	})
}

type ExecuteOrderRequest struct {
	OrderId int64 `json:"order_id"`
}

func (e *Endpoints) ExecuteOrder(c echo.Context) error {
	// this method should do following:
	// lower operator balance on sold currency
	// increase operator balance on topped currency
	// mark operator as free
	// mark transaction as finished

	return nil
}
