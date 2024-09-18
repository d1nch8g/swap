// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"
)

type Balance struct {
	ID         int64   `json:"id"`
	UserID     int64   `json:"user_id"`
	CurrencyID int64   `json:"currency_id"`
	Balance    float64 `json:"balance"`
	Address    string  `json:"address"`
}

type BotMessage struct {
	UserID    *int64    `json:"user_id"`
	OrderID   *int64    `json:"order_id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
	Checked   bool      `json:"checked"`
}

type CardConfirmation struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	CurrencyID int64  `json:"currency_id"`
	Address    string `json:"address"`
	Verified   bool   `json:"verified"`
	Image      []byte `json:"image"`
}

type Currency struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Exchanger struct {
	ID                         int64   `json:"id"`
	Inmin                      float64 `json:"inmin"`
	Description                string  `json:"description"`
	RequirePaymentVerification bool    `json:"require_payment_verification"`
	InCurrency                 int64   `json:"in_currency"`
	OutCurrency                int64   `json:"out_currency"`
}

type Order struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id"`
	OperatorID       int64     `json:"operator_id"`
	ExchangerID      int64     `json:"exchanger_id"`
	AmountIn         float64   `json:"amount_in"`
	AmountOut        float64   `json:"amount_out"`
	ReceiveAddress   string    `json:"receive_address"`
	CreatedAt        time.Time `json:"created_at"`
	Cancelled        bool      `json:"cancelled"`
	Finished         bool      `json:"finished"`
	ConfirmImage     []byte    `json:"confirm_image"`
	PaymentConfirmed bool      `json:"payment_confirmed"`
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
	Passwhash string `json:"passwhash"`
	Admin     bool   `json:"admin"`
	Operator  bool   `json:"operator"`
	Token     string `json:"token"`
	Busy      bool   `json:"busy"`
}
