package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/inswap/gen/database"
)

type Busy struct {
	Busy bool `json:"busy"`
}

// ChangeBusy godoc
//
//	@Summary	Change busy status for admin operator
//	@Param		status	body	Busy	true	"Busy status"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/change-busy [post]
func (e *Endpoints) ChangeBusy(c echo.Context) error {
	var req Busy
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	u, err = e.db.UpdateUserBusy(c.Request().Context(), database.UpdateUserBusyParams{
		Email: u.Email,
		Busy:  req.Busy,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

type Orders struct {
	Orders []database.Order `json:"orders"`
}

type OrderInfo struct {
	Id         int64  `json:"id"`
	CurrencyIn string `json:"currency_in"`
	AmountIn   string
}

// GetOrders godoc
//
//	@Summary	Get active orders bound to specific operator
//	@Success	200	{object}	Orders
//	@Security	ApiKeyAuth
//	@Router		/operator/get-orders [get]
func (e *Endpoints) GetOrders(c echo.Context) error {
	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

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
		Orders: orders,
	})
}

type CreateBalanceRequest struct {
	CurrencyId int64   `json:"currency_id"`
	Balance    float64 `json:"balance"`
	Address    string  `json:"address"`
}

// CreateBalance godoc
//
//	@Summary	Create new operator balance
//	@Param		status	body	CreateBalanceRequest	true	"Create balance parameters"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/create-balance [post]
func (e *Endpoints) CreateBalance(c echo.Context) error {
	var req CreateBalanceRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

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
	BalanceId    int64   `json:"balance_id"`
	CurrencyCode string  `json:"currency_code"`
	Balance      float64 `json:"balance"`
	Address      string  `json:"address"`
}

// UpdateBalance godoc
//
//	@Summary	Update operator currency balance
//	@Param		status	body	UpdateBalanceRequest	true	"Update balance parameters"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/operator/update-balance [post]
func (e *Endpoints) UpdateBalance(c echo.Context) error {
	var req UpdateBalanceRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	if u.Busy {
		c.Response().WriteHeader(http.StatusConflict)
		_, err := c.Response().Write([]byte("unable to update balance while there are unfinished operations"))
		return err
	}

	curr, err := e.db.GetCurrencyByCode(c.Request().Context(), req.CurrencyCode)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	bal, err := e.db.GetBalanceById(c.Request().Context(), database.GetBalanceByIdParams{
		ID:     req.BalanceId,
		UserID: u.ID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			bal, err = e.db.CreateBalance(c.Request().Context(), database.CreateBalanceParams{
				UserID:     u.ID,
				CurrencyID: curr.ID,
				Balance:    req.Balance,
				Address:    req.Address,
			})
			if err != nil {
				c.Response().WriteHeader(http.StatusInternalServerError)
				_, err := c.Response().Write([]byte("unable to access database"))
				return err
			}

		} else {
			c.Response().WriteHeader(http.StatusInternalServerError)
			_, err := c.Response().Write([]byte("unable to access database"))
			return err
		}
	}

	_, err = e.db.UpdateBalance(c.Request().Context(), database.UpdateBalanceParams{
		ID:      bal.ID,
		UserID:  u.ID,
		Balance: req.Balance,
		Address: req.Address,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}
	return nil
}

type Balances struct {
	Balances []Balance `json:"balances"`
}

type Balance struct {
	Id          int64   `json:"id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Balance     float64 `json:"balance"`
}

// ListBalances godoc
//
//	@Summary	List operator currency balances
//	@Success	200 {object}	Balances
//	@Security	ApiKeyAuth
//	@Router		/admin/list-balances [get]
func (e *Endpoints) ListBalances(c echo.Context) error {
	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	dbbalances, err := e.db.ListBalances(c.Request().Context(), u.ID)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	var balances []Balance
	for _, bal := range dbbalances {
		curr, err := e.db.GetCurrencyById(c.Request().Context(), bal.CurrencyID)
		if err != nil {
			c.Response().WriteHeader(http.StatusInternalServerError)
			_, err := c.Response().Write([]byte("unable to access database"))
			return err
		}

		balances = append(balances, Balance{
			Id:          bal.ID,
			Code:        curr.Code,
			Description: curr.Description,
			Address:     bal.Address,
			Balance:     bal.Balance,
		})
	}

	return c.JSON(http.StatusOK, &Balances{
		Balances: balances,
	})
}

type RemoveBalanceRequest struct {
	Id int64 `json:"id"`
}

// RemoveBalance godoc
//
//	@Summary	Remove operators balance
//	@Param		status	body	RemoveBalanceRequest	true	"Balance id"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/remove-balance [delete]
func (e *Endpoints) RemoveBalance(c echo.Context) error {
	var req RemoveBalanceRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	err = e.db.RemoveBalance(c.Request().Context(), database.RemoveBalanceParams{
		ID:     req.Id,
		UserID: u.ID,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

type ExecuteOrderRequest struct {
	OrderId int64 `json:"order_id"`
}

// ExecuteOrder godoc
//
//	@Summary	Execute order and change operator balances, update busy
//	@Param		status	body	ExecuteOrderRequest	true	"Execute order parameters"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/execute-order [post]
func (e *Endpoints) ExecuteOrder(c echo.Context) error {
	// Lower operator balance on sold currency and increase on bought
	var req ExecuteOrderRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

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

	// Send email notification about order being finished.
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

// CancelOrder godoc
//
//	@Summary	Cancel user order and mark it as cancelled
//	@Param		status	body	CancelOrderRequest	true	"Cancel order parameters"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/cancel-order [post]
func (e *Endpoints) CancelOrder(c echo.Context) error {
	var req CancelOrderRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

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

// GetCardConfirmations godoc
//
//	@Summary	Get user credit card approval images with parameters
//	@Success	200 {object}	GetCardConfirmationsResponse
//	@Security	ApiKeyAuth
//	@Router		/admin/get-card-confirmations [get]
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

// ApproveCard godoc
//
//	@Summary	Mark user credit card as approved
//	@Param		status	body	ApproveCardConfirmationRequest	true	"Approve card request"
//	@Success	200
//	@Security	ApiKeyAuth
//	@Router		/admin/approve-card [post]
func (e *Endpoints) ApproveCard(c echo.Context) error {
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
