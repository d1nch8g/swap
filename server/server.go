package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

// add endpoint to give exchangers info for bestchange

func Run(dir, port, tls string, e *echo.Echo, d *database.Queries, b *bestchange.Client) {
	e.Use(middleware.Logger())

	e.Static("/", dir)

	api := e.Group("/api")

	userSvc := &UserService{
		db: d,
		e:  e,
		bc: b,
	}

	api.POST("/login", userSvc.Login)
	api.POST("/create", userSvc.CreateUser)

	admin := api.Group("/admin")

	admin.GET("/getorders", userSvc.GetOrders)
	admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		token := c.Request().Header["Token"]
		if token == nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return false, nil
		}

		_, err := d.GetUserByToken(c.Request().Context(), token[0])
		if err != nil {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return false, nil
		}
		return true, nil
	}))

	if tls != "" {
		e.Logger.Fatal(e.StartAutoTLS(tls))
	}
	e.Logger.Fatal(e.Start("localhost:" + port))
}
