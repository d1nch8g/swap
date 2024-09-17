package server

import (
	"net/http"

	"github.com/d1nch8g/swap/bestchange"
	"github.com/d1nch8g/swap/email"
	"github.com/d1nch8g/swap/gen/database"
	"github.com/d1nch8g/swap/gen/web"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Endpoints struct {
	db       *database.Queries
	e        *echo.Echo
	bc       *bestchange.Client
	mail     *email.Mailer
	pgx      *pgxpool.Pool
	host     string
	email    string
	telegram string
}

func Run(port, host, certFile, keyFile, email, telegram string, e *echo.Echo, p *pgxpool.Pool, d *database.Queries, b *bestchange.Client, mail *email.Mailer) {
	endpoints := &Endpoints{
		db:       d,
		e:        e,
		bc:       b,
		mail:     mail,
		pgx:      p,
		host:     host,
		email:    email,
		telegram: telegram,
	}

	e.Use(middleware.Logger())

	staticDir := []string{
		"/",
		"/contacts",
		"/login",
		"/register",
		"/rules",
		"/profile",
		"/operator",
		"/currencies",
		"/exchangers",
		"/card-confirmations",
		"/transfer",
		"/order",
		"/validate-card",
	}

	for _, path := range staticDir {
		handler := http.FileServer(web.AssetFile())
		e.GET(path, echo.WrapHandler(http.StripPrefix(path, handler)))
	}
	for _, path := range web.AssetNames() {
		handler := http.FileServer(web.AssetFile())
		e.GET(path, echo.WrapHandler(handler))
	}

	api := e.Group("/api")

	api.GET("/info", endpoints.Info)
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
	user.GET("/self-info", endpoints.SelfInfo)

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
	operator.GET("/card-confirmations", endpoints.CardConfirmations)

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

	if keyFile != "" && certFile != "" {
		e.Logger.Fatal(e.StartTLS(":"+port, certFile, keyFile))
	}
	e.Logger.Fatal(e.Start(":" + port))
}
