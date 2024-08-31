package server

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
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

func (s *UserService) Login(c echo.Context) error {
	email := c.Request().Header["email"]
	password := c.Request().Header["password"]

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
