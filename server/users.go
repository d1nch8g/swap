package server

import (
	"crypto/sha512"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *Endpoints) CreateUser(c echo.Context) error {
	var createUser CreateUserRequest
	err := c.Bind(&createUser)
	if err != nil {
		c.Response().WriteHeader(http.StatusForbidden)
		_, err := c.Response().Write([]byte("unable to bind user request"))
		return err
	}

	uuid := uuid.New().String()

	err = e.mail.UserVerifyEmail(createUser.Email, uuid)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable send email notification"))
		return err
	}

	hasher := sha512.New()
	hasher.Write([]byte(createUser.Password))
	passhash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	_, err = e.db.CreateUser(c.Request().Context(), database.CreateUserParams{
		Email:     createUser.Email,
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

type CreateOrderRequest struct {
	Email  string  `json:"email"`
	Input  string  `json:"input"`
	Ouput  string  `json:"output"`
	Amount float64 `json:"amount"`
}

func (e *Endpoints) CreateOrder(c echo.Context) error {

	// Estimate exchange rate for order and create it for user with one free admin
	return nil
}

// func (m *Endpoints)

// verify card
// payment approve
