package server

import (
	"net/http"
	"strings"

	"github.com/d1nch8g/swap/gen/database"
	"github.com/labstack/echo/v4"
)

// CreateCurrency godoc
//
//	@Summary	Create new currency in exchanger
//	@Param		status	body	database.CreateCurrencyParams	true	"Create currency params"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/create-currency [post]
func (e *Endpoints) CreateCurrency(c echo.Context) error {
	var req database.CreateCurrencyParams
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	_, err = e.db.CreateCurrency(c.Request().Context(), req)
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

// RemoveCurrency godoc
//
//	@Summary	Remove currency from currency list
//	@Param		status	body	RemoveCurrencyRequest	true	"Remove currency parameter"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/remove-currency [delete]
func (e *Endpoints) RemoveCurrency(c echo.Context) error {
	var req RemoveCurrencyRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	err = e.db.RemoveCurrency(c.Request().Context(), req.Code)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

type CreateExchangerRequest struct {
	Description        string  `json:"description"`
	Inmin              float64 `json:"inmin"`
	PaymentVerfication bool    `json:"payment_verification"`
	InCurrency         string  `json:"in_currency"`
	OutCurrency        string  `json:"out_currency"`
}

// CreateExchanger godoc
//
//	@Summary	Create new exchanger with provided currencies
//	@Param		status	body	CreateExchangerRequest	true	"Create exchanger parameters"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/create-exchanger [post]
func (e *Endpoints) CreateExchanger(c echo.Context) error {
	var req CreateExchangerRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	inCurr, err := e.db.GetCurrencyByCode(c.Request().Context(), req.InCurrency)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	outCurr, err := e.db.GetCurrencyByCode(c.Request().Context(), req.OutCurrency)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	_, err = e.db.CreateExchanger(c.Request().Context(), database.CreateExchangerParams{
		Description:                req.Description,
		Inmin:                      req.Inmin,
		RequirePaymentVerification: req.PaymentVerfication,
		InCurrency:                 inCurr.ID,
		OutCurrency:                outCurr.ID,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

type RemoveExchangerRequest struct {
	Id int64 `json:"id"`
}

// RemoveExchanger godoc
//
//	@Summary	Remove existing exchanger from API
//	@Param		status	body	RemoveExchangerRequest	true	"Remove exchanger parameters"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/remove-exchanger [delete]
func (e *Endpoints) RemoveExchanger(c echo.Context) error {
	var req RemoveExchangerRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	err = e.db.RemoveExchanger(c.Request().Context(), req.Id)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

// CheckIfAdmin godoc
//
//	@Summary	Check if user is an admin
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/check-if-admin [post]
func (e *Endpoints) CheckIfAdmin(c echo.Context) error {
	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	if !u.Admin {
		c.Response().WriteHeader(http.StatusConflict)
		_, err := c.Response().Write([]byte("not an admin"))
		return err
	}

	return nil
}
