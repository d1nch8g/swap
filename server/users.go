package server

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/inswap/gen/database"
)

// Login godoc
//
//	@Summary	Login and get auth key
//	@Param		email		header		string	true	"Email"
//	@Param		password	header		string	true	"Password"
//	@Success	200			{string}	string	token
//	@Router		/login [post]
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

	if user.Token != "nil" {
		_, err = c.Response().Write([]byte(user.Token))
		return err
	}

	token := uuid.New().String()
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

// ListOrders godoc
//
//	@Summary	List user's orders
//	@Success	200			{object}	Orders	orders
//	@Security	ApiKeyAuth
//	@Router		/user/list-orders [post]
func (e *Endpoints) ListOrders(c echo.Context) error {
	token := strings.ReplaceAll(c.Request().Header["Authorization"][0], "Bearer ", "")

	u, err := e.db.GetUserByToken(c.Request().Context(), token)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	orders, err := e.db.GetOrdersForUser(c.Request().Context(), u.ID)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &Orders{
		Orders: orders,
	})
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser godoc
//
//	@Summary	Create new user request
//	@Param		status	body	CreateUserRequest	true	"Create user request"
//	@Success	200
//	@Router		/create-user [post]
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

	u, err := e.db.CreateUser(c.Request().Context(), database.CreateUserParams{
		Email:     req.Email,
		Passwhash: passhash,
		Token:     uuid,
		Verified:  false,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint ") {
			u, err = e.db.GetUser(c.Request().Context(), req.Email)
			if err != nil {
				c.Response().WriteHeader(http.StatusInternalServerError)
				_, err := c.Response().Write([]byte("unable to create new user"))
				return err
			}
			if !u.Verified {
				_, err := e.db.UpdateUserTokenAndPassHash(c.Request().Context(), database.UpdateUserTokenAndPassHashParams{
					Email:     u.Email,
					Token:     uuid,
					Passwhash: passhash,
				})
				if err != nil {
					c.Response().WriteHeader(http.StatusInternalServerError)
					_, err := c.Response().Write([]byte("unable to create new user"))
					return err
				}
				return nil
			}
		}
		c.Response().WriteHeader(http.StatusForbidden)
		_, err := c.Response().Write([]byte("unable to create new user"))
		return err
	}

	return nil
}

// VerifyEmail godoc
//
//	@Summary	Verify user email address
//	@Param		uuid	path	string	true	"UUID sent by email"
//	@Success	200
//	@Router		/verify/{uuid} [get]
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

	return c.String(http.StatusCreated, "email have been verified")
}

type Currencies struct {
	Currencies []database.Currency `json:"currencies"`
}

// ListCurrencies godoc
//
//	@Summary	Verify user email address
//	@Success	200	{object}	Currencies
//	@Router		/list-currencies [get]
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

type CurrentRateResponse struct {
	Amount float64 `json:"amount"`
}

// CurrentRate godoc
//
//	@Summary	Current rate at specific currency
//	@Param		currency_in		path		string	true	"Currency in"
//	@Param		currency_out	path		string	true	"Currency out"
//	@Param		amount			path		int		true	"Amount in"
//	@Success	200				{object}	CurrentRateResponse
//	@Router		/current-rate [get]
func (e *Endpoints) CurrentRate(c echo.Context) error {
	currencyIn := c.QueryParam("currency_in")
	currencyOut := c.QueryParam("currency_out")
	amountString := c.QueryParam("amount")
	amount, err := strconv.ParseFloat(amountString, 64)

	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to parse amount"))
		return err
	}

	currIn, err := e.db.GetCurrencyByCode(c.Request().Context(), currencyIn)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	currOut, err := e.db.GetCurrencyByCode(c.Request().Context(), currencyOut)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	exch, err := e.db.GetExchangerByCurrencyIds(c.Request().Context(), database.GetExchangerByCurrencyIdsParams{
		InCurrency:  currIn.ID,
		OutCurrency: currOut.ID,
	})
	if err != nil {
		if err.Error() == "no rows in result set" {
			c.Response().WriteHeader(http.StatusForbidden)
			_, err := c.Response().Write([]byte("selected exchangers pair does not exist"))
			return err
		}
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}
	if amount < exch.Inmin {
		c.Response().WriteHeader(http.StatusConflict)
		_, err := c.Response().Write([]byte(fmt.Sprintf("not over minimum operation %f", exch.Inmin)))
		return err
	}

	rez, err := e.bc.EstimateOperation(currencyIn, currencyOut)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access bestchange API"))
		return err
	}

	return c.JSON(http.StatusOK, &CurrentRateResponse{
		Amount: amount / rez,
	})
}

type Exchangers struct {
	Exchangers []database.Exchanger `json:"exchangers"`
}

// ListExchangers godoc
//
//	@Summary	List existing exchangers
//	@Success	200	{object}	Exchangers
//	@Router		/list-exchangers [get]
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
	Email       string  `json:"email"`
	InCurrency  string  `json:"in_currency"`
	OutCurrency string  `json:"out_currency"`
	Amount      float64 `json:"amount"`
	Address     string  `json:"address"`
}

type CreateOrderResponse struct {
	InAmount        float64 `json:"in_amount"`
	OutAmount       float64 `json:"out_amount"`
	TransferAddress string  `json:"transfer_address"`
	OrderNumber     int64   `json:"order_number"`
}

// ListExchangers godoc
//
//	@Summary	Create order to exchange specific currency
//	@Param		status	body		CreateOrderRequest	true	"Request parameters"
//	@Success	200		{object}	CreateOrderResponse
//	@Router		/create-order [post]
func (e *Endpoints) CreateOrder(c echo.Context) error {
	var req CreateOrderRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	// Check if input is over minimum for given exchanger
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

	exch, err := e.db.GetExchangerByCurrencyIds(c.Request().Context(), database.GetExchangerByCurrencyIdsParams{
		InCurrency:  inCurr.ID,
		OutCurrency: outCurr.ID,
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
	u, err := e.db.GetUser(c.Request().Context(), req.Email)
	if err != nil {
		u, err = e.db.CreateUser(c.Request().Context(), database.CreateUserParams{
			Email:     req.Email,
			Passwhash: "nil",
			Token:     "nil",
		})
		if err != nil {
			c.Response().WriteHeader(http.StatusInternalServerError)
			_, err := c.Response().Write([]byte("unable to access database"))
			return err
		}
	}

	if exch.RequirePaymentVerification {
		pc, err := e.db.GetCardConfirmation(c.Request().Context(), database.GetCardConfirmationParams{
			UserID:     u.ID,
			CurrencyID: inCurr.ID,
		})
		if err != nil {
			_, err := e.db.CreateCardConfirmation(c.Request().Context(), database.CreateCardConfirmationParams{
				UserID:     u.ID,
				CurrencyID: inCurr.ID,
				Address:    req.Address,
				Verified:   false,
			})
			if err != nil {
				c.Response().WriteHeader(http.StatusInternalServerError)
				_, err := c.Response().Write([]byte("unable to access database or exchanger does not exist"))
				return err
			}
			return c.String(http.StatusForbidden, "required card confirmation")
		}
		if !pc.Verified {
			return c.String(http.StatusForbidden, "required card confirmation")
		}
	}

	// Estimate exchange rate based on bestchange API and calculate output
	rate, err := e.bc.EstimateOperation(inCurr.Code, outCurr.Code)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access bestchange API"))
		return err
	}

	outAmount := req.Amount / rate

	// Find free operator with proper calculated amount for given currency
	admins, err := e.db.GetFreeAdmins(c.Request().Context())
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}
	var operator *database.User
	var addr string
	for _, admin := range admins {
		balances, err := e.db.ListBalances(c.Request().Context(), admin.ID)
		if err != nil {
			c.Response().WriteHeader(http.StatusInternalServerError)
			_, err := c.Response().Write([]byte("unable to access database"))
			return err
		}
		var goodOutBalance bool
		var existInBalance bool
		for _, balance := range balances {
			if balance.CurrencyID == outCurr.ID && balance.Balance > outAmount {
				operator = &admin
				goodOutBalance = true
			}
			if balance.CurrencyID == inCurr.ID {
				addr = balance.Address
				existInBalance = true
			}
		}
		if goodOutBalance && existInBalance {
			break
		}
	}

	if operator == nil || addr == "" {
		c.Response().WriteHeader(http.StatusConflict)
		_, err := c.Response().Write([]byte("all operators are busy"))
		return err
	}

	// Mark operator as busy and create new order
	_, err = e.db.UpdateUserBusy(c.Request().Context(), database.UpdateUserBusyParams{
		Email: operator.Email,
		Busy:  true,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	order, err := e.db.CreateOrder(c.Request().Context(), database.CreateOrderParams{
		UserID:         u.ID,
		OperatorID:     operator.ID,
		ExchangerID:    exch.ID,
		AmountIn:       req.Amount,
		AmountOut:      outAmount,
		ReceiveAddress: req.Address,
		Cancelled:      false,
		Finished:       false,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	// Send email notification
	err = e.mail.OrderCreated(req.Email, req.InCurrency, req.OutCurrency, fmt.Sprintf("%f", req.Amount), req.Address)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to send email notification"))
		return err
	}

	return c.JSON(http.StatusOK, &CreateOrderResponse{
		InAmount:        req.Amount,
		OutAmount:       outAmount,
		TransferAddress: addr,
		OrderNumber:     order.ID,
	})
}

type ValidateCardRequest struct {
	Email      string `json:"email"`
	CurrencyId int64  `json:"currency_id"`
}

func (e *Endpoints) ValidateCard(c echo.Context) error {
	var req ValidateCardRequest
	err := c.Bind(&req)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		_, err := c.Response().Write([]byte("unable to unmarshal request"))
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to read file"))
		return err
	}
	src, err := file.Open()
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to read file"))
		return err
	}

	img, err := io.ReadAll(src)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to read file"))
		return err
	}

	u, err := e.db.GetUser(c.Request().Context(), req.Email)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to read file"))
		return err
	}

	cc, err := e.db.GetCardConfirmation(c.Request().Context(), database.GetCardConfirmationParams{
		UserID:     u.ID,
		CurrencyID: req.CurrencyId,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	_, err = e.db.UpdateCardConfirmationImage(c.Request().Context(), database.UpdateCardConfirmationImageParams{
		ID:    cc.ID,
		Image: img,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

// approve payment operated
