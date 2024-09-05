package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ion.lc/d1nhc8g/inswap/bestchange"
	"ion.lc/d1nhc8g/inswap/email"
	"ion.lc/d1nhc8g/inswap/gen/database"
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

	api.POST("/create-user", endpoints.CreateUser)
	api.POST("/create-order", endpoints.CreateOrder)
	api.POST("/validate-card", endpoints.ValidateCard)
	api.GET("/verify/:uuid", endpoints.VerifyEmail)
	api.GET("/list-currencies", endpoints.ListCurrencies)
	api.GET("/list-exchangers", endpoints.ListExchangers)
	api.POST("/login", endpoints.Login)

	admin := api.Group("/admin", middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
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

		if !u.Verified {
			c.Response().WriteHeader(http.StatusUnauthorized)
			_, err = c.Response().Write([]byte("verify user email first"))
			return false, err
		}

		if err != nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return false, nil
		}
		return true, nil
	}))

	admin.GET("/get-orders", endpoints.GetOrders)
	admin.POST("/create-currency", endpoints.CreateCurrency)
	admin.POST("/remove-currency", endpoints.RemoveCurrency)
	admin.POST("/create-exchanger", endpoints.CreateExchanger)
	admin.POST("/remove-exchanger", endpoints.RemoveExchanger)
	admin.POST("/create-balance", endpoints.CreateBalance)
	admin.POST("/list-balances", endpoints.ListBalances)
	admin.POST("/update-balance", endpoints.UpdateBalance)
	admin.POST("/execute-order", endpoints.ExecuteOrder)
	admin.POST("/get-card-confirmations", endpoints.GetCardConfirmations)
	admin.POST("/approve-card", endpoints.ApproveCardConfirmation)

	if tls != "" {
		e.Logger.Fatal(e.StartAutoTLS(tls))
	}
	e.Logger.Fatal(e.Start("localhost:" + port))
}
