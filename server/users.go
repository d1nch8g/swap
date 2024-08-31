package server

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type UserService struct {
	db *database.Queries
	e  *echo.Echo
	bc *bestchange.Client
}

// This method will issue token for user and give it to new users.
func (s *UserService) Login(c echo.Context) error {
	email := c.Request().Header["Email"]
	password := c.Request().Header["Password"]

	if email == nil || password == nil {
		c.Response().WriteHeader(http.StatusUnauthorized)
		_, err := c.Response().Write([]byte("empty login or password"))
		return err
	}

	user, err := s.db.GetUser(c.Request().Context(), email[0])
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

	_, err = s.db.UpdateUserToken(c.Request().Context(), database.UpdateUserTokenParams{
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

type Orders struct {
	ActiveOrders []database.Order `json:"orders"`
}

func (s *UserService) GetOrders(c echo.Context) error {
	orders, err := s.db.OrdersUnfinished(c.Request().Context())
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, err := c.Response().Write([]byte("unable to access database"))
		return err
	}

	return c.JSON(http.StatusOK, &Orders{
		ActiveOrders: orders,
	})
}

type CreateUserRequest struct {
	Email string `json:"email"`
}

func (s *UserService) CreateUser(c echo.Context) error {
	var createUser CreateUserRequest
	err := c.Bind(&createUser)
	if err != nil {
		c.Response().WriteHeader(http.StatusForbidden)
		_, err := c.Response().Write([]byte("unable to bind user request"))
		return err
	}

	_, err = s.db.CreateUser(c.Request().Context(), database.CreateUserParams{
		Email:     createUser.Email,
		Passwhash: "nil",
		Token:     "nil",
	})
	if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint ") {
		c.Response().WriteHeader(http.StatusForbidden)
		_, err := c.Response().Write([]byte("unable to create new user"))
		return err
	}

	return nil
}
