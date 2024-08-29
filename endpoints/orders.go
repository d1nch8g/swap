package endpoints

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"ion.lc/d1nhc8g/bitchange/bestchange"
)

func (m *mapper) ActualParams(c echo.Context) error {
	var request struct {
		Give    uint16 `query:"give"`
		Receive uint16 `query:"receive"`
	}
	c.Bind(&request)

	sbpton, err := m.bc.Rates(request.Give, request.Receive)
	if err != nil {
		panic(err)
	}
	fmt.Println(" ===== SBT-TON ===== Give SBT, receive TON")
	bestchange.PrintTable(sbpton)

	tonsbp, err := m.bc.Rates(request.Receive, request.Give)
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

func (m *mapper) CreateOrder(c echo.Context) error {
	return nil
}
