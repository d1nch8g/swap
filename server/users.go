package server

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
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
			if u.Verified == false {
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
	Email   string  `json:"email"`
	Input   string  `json:"input"`
	Ouput   string  `json:"output"`
	Amount  float64 `json:"amount"`
	Address string  `json:"address"`
}

type CreateOrderResponse struct {
	InAmount        float64 `json:"in_amount"`
	OutAmount       float64 `json:"out_amount"`
	TransferAddress string  `json:"transfer_address"`
}

func (e *Endpoints) CreateOrder(c echo.Context) error {
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
			return c.String(http.StatusConflict, "required card confirmation")
		}
		if !pc.Verified {
			return c.String(http.StatusConflict, "required card confirmation")
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
		for _, balance := range balances {
			if balance.CurrencyID == outCurr.ID && balance.Balance > outAmount {
				addr = balance.Address
				operator = &admin
				break
			}
		}
	}

	if operator == nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
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

	_, err = e.db.CreateOrder(c.Request().Context(), database.CreateOrderParams{
		UserID:      u.ID,
		OperatorID:  operator.ID,
		ExchangerID: exch.ID,
		AmountIn:    req.Amount,
		AmountOut:   outAmount,
		Finished:    false,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	// Send email notification
	err = e.mail.OrderCreated(req.Email, req.Input, req.Ouput, fmt.Sprintf("%d", req.Amount), req.Address)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to send email notification"))
		return err
	}

	return c.JSON(http.StatusOK, &CreateOrderResponse{
		InAmount:        req.Amount,
		OutAmount:       outAmount,
		TransferAddress: addr,
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
		ID:       cc.ID,
		Image:    img,
		Verified: false,
	})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return nil
}

// approve payment operated
