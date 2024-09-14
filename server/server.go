package server

import (
	"net/http"
	"path"

	"github.com/jackc/pgx/v5/pgxpool"
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
	pgx  *pgxpool.Pool
}

func Run(dir, port, host, certDir string, e *echo.Echo, p *pgxpool.Pool, d *database.Queries, b *bestchange.Client, mail *email.Mailer) {
	endpoints := &Endpoints{
		db:   d,
		e:    e,
		bc:   b,
		pgx:  p,
		mail: mail,
	}

	e.Use(middleware.Logger())

	e.Static("/", dir)
	e.Static("/contacts", dir)
	e.Static("/login", dir)
	e.Static("/register", dir)
	e.Static("/rules", dir)
	e.Static("/profile", dir)
	e.Static("/operator", dir)
	e.Static("/currencies", dir)
	e.Static("/exchangers", dir)
	e.Static("/operators", dir)
	e.Static("/transfer", dir)
	e.Static("/order", dir)
	e.Static("/validate-card", dir)

	api := e.Group("/api")

	api.POST("/create-user", endpoints.CreateUser)
	api.POST("/create-order", endpoints.CreateOrder)
	api.POST("/validate-card", endpoints.ValidateCard)
	api.POST("/confirm-payment", endpoints.ConfirmPayment)
	api.GET("/verify/:uuid", endpoints.VerifyEmail)
	api.GET("/list-currencies", endpoints.ListCurrencies)
	api.GET("/list-exchangers", endpoints.ListExchangers)
	api.GET("/current-rate", endpoints.CurrentRate)
	api.GET("/order-status", endpoints.OrderStatus)
	api.POST("/login", endpoints.Login)

	user := api.Group("/user", middleware.KeyAuth(func(auth string, c echo.Context) (bool, error) {
		u, err := d.GetUserByToken(c.Request().Context(), auth)
		if err != nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			_, err = c.Response().Write([]byte("user is not found"))
			return false, err
		}

		if !u.Verified {
			c.Response().WriteHeader(http.StatusUnauthorized)
			_, err = c.Response().Write([]byte("verify user email first"))
			return false, err
		}

		return true, nil
	}))

	user.GET("/list-orders", endpoints.ListOrders)

	operator := api.Group("/operator", middleware.KeyAuth(func(auth string, c echo.Context) (bool, error) {
		u, err := d.GetUserByToken(c.Request().Context(), auth)
		if err != nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			_, err = c.Response().Write([]byte("user is not found"))
			return false, err
		}

		if !u.Operator {
			c.Response().WriteHeader(http.StatusUnauthorized)
			_, err = c.Response().Write([]byte("user is not an operator"))
			return false, err
		}

		return true, nil
	}))

	operator.POST("/change-busy", endpoints.ChangeBusy)
	operator.GET("/get-orders", endpoints.GetOrders)
	operator.GET("/finished-orders", endpoints.FinishedOrders)
	operator.POST("/create-balance", endpoints.CreateBalance)
	operator.GET("/list-balances", endpoints.ListBalances)
	operator.POST("/update-balance", endpoints.UpdateBalance)
	operator.DELETE("/remove-balance", endpoints.RemoveBalance)
	operator.POST("/execute-order", endpoints.ExecuteOrder)
	operator.POST("/cancel-order", endpoints.CancelOrder)
	operator.GET("/get-card-confirmations", endpoints.GetCardConfirmations)
	operator.POST("/approve-card", endpoints.ApproveCard)
	operator.DELETE("/cancel-card", endpoints.CancelCard)

	admin := api.Group("/admin", middleware.KeyAuth(func(auth string, c echo.Context) (bool, error) {
		u, err := d.GetUserByToken(c.Request().Context(), auth)
		if err != nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			_, err = c.Response().Write([]byte("user is not found"))
			return false, err
		}

		if !u.Admin {
			c.Response().WriteHeader(http.StatusUnauthorized)
			_, err = c.Response().Write([]byte("user is not an admin"))
			return false, err
		}

		return true, nil
	}))

	admin.POST("/check-if-admin", endpoints.CheckIfAdmin)
	admin.POST("/create-currency", endpoints.CreateCurrency)
	admin.DELETE("/remove-currency", endpoints.RemoveCurrency)
	admin.POST("/create-exchanger", endpoints.CreateExchanger)
	admin.DELETE("/remove-exchanger", endpoints.RemoveExchanger)

	if certDir != "" {
		e.Logger.Fatal(e.StartTLS(host+":"+port, path.Join(certDir, host+".crt"), path.Join(certDir, host+".key")))
	}
	e.Logger.Fatal(e.Start(host + ":" + port))
}
