package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/email"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type Endpoints struct {
	db   *database.Queries
	e    *echo.Echo
	bc   *bestchange.Client
	mail *email.Mailer
}

func Run(dir, port, tls string, e *echo.Echo, d *database.Queries, b *bestchange.Client, mail *email.Mailer) {
	endpoints := &Endpoints{
		db:   d,
		e:    e,
		bc:   b,
		mail: mail,
	}

	e.Use(middleware.Logger())

	e.Static("/", dir)

	api := e.Group("/api")

	api.POST("/createuser", endpoints.CreateUser)
	api.POST("/createorder", endpoints.CreateOrder)

	api.POST("/login", endpoints.Login)
	admin := api.Group("/admin")

	admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		token := c.Request().Header["Token"]
		if token == nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return false, nil
		}

		u, err := d.GetUserByToken(c.Request().Context(), token[0])
		if !u.Admin {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return false, nil
		}

		if err != nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return false, nil
		}
		return true, nil
	}))
	admin.GET("/getorders", endpoints.GetOrders)
	admin.POST("/createcurrency", endpoints.CreateCurrency)

	if tls != "" {
		e.Logger.Fatal(e.StartAutoTLS(tls))
	}
	e.Logger.Fatal(e.Start("localhost:" + port))
}
