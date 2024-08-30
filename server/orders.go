package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
)

type orderservice struct {
	db *database.Queries
	e  *echo.Echo
	bc *bestchange.Client
}

func (m *orderservice) ActualParams(c echo.Context) error {
	give := c.Param("give")
	receive := c.Param("receive")

	fmt.Println("taking rates")
	sbpton, err := m.bc.Rates(give, receive)
	if err != nil {
		panic(err)
	}
	fmt.Println(" ===== SBT-TON ===== Give SBT, receive TON")
	bestchange.PrintTable(sbpton)

	fmt.Println("taking rates second")

	tonsbp, err := m.bc.Rates(receive, give)
	if err != nil {
		panic(err)
	}
	fmt.Println(" ===== TON-SBP ===== Give TON, receive SBT")
	bestchange.PrintTable(tonsbp)

	buy, sell, avg := m.bc.EstimateRates(sbpton, tonsbp)
	fmt.Printf("Buy: %f, Sell: %f, Avg: %f\n", avg, buy, sell)
	// min max exchange rate
	return nil
}

func (m *orderservice) CreateOrder(c echo.Context) error {
	return nil
}
