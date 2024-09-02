package server

import (
	"crypto/sha512"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/inswap/gen/database"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *Endpoints) CreateUser(c echo.Context) error {
	var req CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	uuid := uuid.New().String()

	err = e.mail.UserVerifyEmail(req.Email, uuid)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable send email notification"))
		return err
	}

	hasher := sha512.New()
	hasher.Write([]byte(req.Password))
	passhash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	_, err = e.db.CreateUser(c.Request().Context(), database.CreateUserParams{
		Email:     req.Email,
		Passwhash: passhash,
		Token:     uuid,
		Verified:  false,
	})
	if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint ") {
		c.Response().WriteHeader(http.StatusForbidden)
		_, err := c.Response().Write([]byte("unable to create new user"))
		return err
	}
	return nil
}

func (e *Endpoints) VerifyEmail(c echo.Context) error {
	u, err := e.db.GetUserByToken(c.Request().Context(), c.Param("uuid"))
	if err != nil {
		c.Response().WriteHeader(http.StatusUnauthorized)
		_, err := c.Response().Write([]byte("unable to get user access"))
		return err
	}

	_, err = e.db.UpdateUserVerified(c.Request().Context(), database.UpdateUserVerifiedParams{
		Email:    u.Email,
		Verified: true,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to get user access"))
		return err
	}
	c.Response().WriteHeader(http.StatusCreated)
	return nil
}

type Currencies struct {
	Currencies []database.Currency `json:"currencies"`
}

func (e *Endpoints) ListCurrencies(c echo.Context) error {
	currs, err := e.db.ListCurrencies(c.Request().Context())
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &Currencies{
		Currencies: currs,
	})
}

type Exchangers struct {
	Exchangers []database.Exchanger `json:"exchangers"`
}

func (e *Endpoints) ListExchangers(c echo.Context) error {
	exchangers, err := e.db.ListExchangers(c.Request().Context())
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &Exchangers{
		Exchangers: exchangers,
	})
}

type CreateOrderRequest struct {
	Email  string  `json:"email"`
	Input  string  `json:"input"`
	Ouput  string  `json:"output"`
	Amount float64 `json:"amount"`
}

func (e *Endpoints) CreateOrder(c echo.Context) error {
	// This method should do following:
	// estimate exchange rate based on bestchange API
	// calculate output amount
	// find free operator with proper calculated amount for given currency
	// if user email not exists create and bind to request
	// mark operator as busy
	// create new order

	var req CreateOrderRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	// Check if input is over minimum for given exchanger
	inCurr, err := e.db.GetCurrencyByCode(c.Request().Context(), req.Input)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	outCurr, err := e.db.GetCurrencyByCode(c.Request().Context(), req.Ouput)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	exch, err := e.db.GetExchangerByCurrencyIds(c.Request().Context(), database.GetExchangerByCurrencyIdsParams{
		Input:  inCurr.ID,
		Output: outCurr.ID,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database or exchanger does not exist"))
		return err
	}

	if req.Amount < exch.Inmin {
		c.Response().WriteHeader(http.StatusConflict)
		_, err := c.Response().Write([]byte("input is too small"))
		return err
	}

	// Check if payment requires verification and check if user payment method is validated
	if exch.RequirePaymentVerification {
		
	}

	return nil
}

// verify card
// approve payment operated
