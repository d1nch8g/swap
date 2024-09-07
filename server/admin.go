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

// @Summary		Login
// @Description	get auth key
// @Produce		json
// @Param		email	header	string	true	"Account ID"
// @Param		password		header	string		true	"Account ID"
// @Success		200		{string} string token
// @Router		/login [post]
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
	// Lower operator balance on sold currency and increase on bought
	var req ExecuteOrderRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := c.Request().Header["Token"][0]

	operator, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	balances, err := e.db.ListBalances(c.Request().Context(), operator.ID)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	order, err := e.db.GetOrder(c.Request().Context(), req.OrderId)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	for _, balance := range balances {
		exch, err := e.db.GetExchangerById(c.Request().Context(), order.ExchangerID)
		if err != nil {
			c.Response().WriteHeader(http.StatusInternalServerError)
			_, err := c.Response().Write([]byte("unable to access database"))
			return err
		}

		if balance.CurrencyID == exch.InCurrency {
			_, err = e.db.UpdateBalance(c.Request().Context(), database.UpdateBalanceParams{
				ID:      balance.CurrencyID,
				UserID:  operator.ID,
				Balance: balance.Balance + order.AmountIn,
			})
			if err != nil {
				c.Response().WriteHeader(http.StatusInternalServerError)
				_, err := c.Response().Write([]byte("unable to access database"))
				return err
			}
		}

		if balance.CurrencyID == exch.OutCurrency {
			_, err = e.db.UpdateBalance(c.Request().Context(), database.UpdateBalanceParams{
				ID:      balance.CurrencyID,
				UserID:  operator.ID,
				Balance: balance.Balance - order.AmountOut,
			})
			if err != nil {
				c.Response().WriteHeader(http.StatusInternalServerError)
				_, err := c.Response().Write([]byte("unable to access database"))
				return err
			}
		}
	}

	// Mark operator as free.
	_, err = e.db.UpdateUserBusy(c.Request().Context(), database.UpdateUserBusyParams{
		Email: operator.Email,
		Busy:  false,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	// Mark transaction as finished.
	_, err = e.db.UpdateOrderFinished(c.Request().Context(), order.ID)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	u, err := e.db.GetUserById(c.Request().Context(), order.UserID)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	exch, err := e.db.GetExchangerById(c.Request().Context(), order.ExchangerID)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	outCurr, err := e.db.GetCurrencyById(c.Request().Context(), exch.OutCurrency)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	err = e.mail.OrderFinished(u.Email, fmt.Sprintf("%f", order.AmountOut), outCurr.Code, order.ReceiveAddress)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to send email notification"))
		return err
	}

	return nil
}

type CancelOrderRequest struct {
	OrderId int64 `json:"order_id"`
}

func (e *Endpoints) CancelOrder(c echo.Context) error {
	var req CancelOrderRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := c.Request().Header["Token"][0]

	operator, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	_, err = e.db.UpdateOrderCancelled(c.Request().Context(), req.OrderId)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	_, err = e.db.UpdateUserBusy(c.Request().Context(), database.UpdateUserBusyParams{
		Email: operator.Email,
		Busy:  false,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}
	return nil

}

type GetCardConfirmationsResponse struct {
	CardConfirmations []database.CardConfirmation `json:"card_confirmations"`
}

func (e *Endpoints) GetCardConfirmations(c echo.Context) error {
	cc, err := e.db.GetCardConfirmations(c.Request().Context())
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &GetCardConfirmationsResponse{
		CardConfirmations: cc,
	})
}

type ApproveCardConfirmationRequest struct {
	ConfirmationId int64 `json:"confirmation_id"`
}

func (e *Endpoints) ApproveCardConfirmation(c echo.Context) error {
	var req ApproveCardConfirmationRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	_, err = e.db.UpdateCardConfirmationVerified(c.Request().Context(), database.UpdateCardConfirmationVerifiedParams{
		ID:       req.ConfirmationId,
		Verified: true,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}
