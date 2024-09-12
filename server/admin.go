package server

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/inswap/gen/database"
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

// CreateExchanger godoc
//
//	@Summary	Create new exchanger with provided currencies
//	@Param		status	body	database.CreateExchangerParams	true	"Create exchanger parameters"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/create-exchanger [post]
func (e *Endpoints) CreateExchanger(c echo.Context) error {
	var req database.CreateExchangerParams
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	_, err = e.db.CreateExchanger(c.Request().Context(), req)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

// RemoveExchanger godoc
//
//	@Summary	Remove existing exchanger from API
//	@Param		status	body	database.RemoveExchangerParams	true	"Remove exchanger parameters"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/remove-exchanger [delete]
func (e *Endpoints) RemoveExchanger(c echo.Context) error {
	var req database.RemoveExchangerParams
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	err = e.db.RemoveExchanger(c.Request().Context(), req)
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
//	@Router		/admin/check-if-admin [delete]
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
