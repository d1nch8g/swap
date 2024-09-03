// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
)

const createBalance = `-- name: CreateBalance :one
INSERT INTO balances (user_id, currency_id, balance, address)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, currency_id, balance, address
`

type CreateBalanceParams struct {
	UserID     int64   `json:"user_id"`
	CurrencyID int64   `json:"currency_id"`
	Balance    float64 `json:"balance"`
	Address    string  `json:"address"`
}

func (q *Queries) CreateBalance(ctx context.Context, arg CreateBalanceParams) (Balance, error) {
	row := q.db.QueryRow(ctx, createBalance,
		arg.UserID,
		arg.CurrencyID,
		arg.Balance,
		arg.Address,
	)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CurrencyID,
		&i.Balance,
		&i.Address,
	)
	return i, err
}

const createCardConfirmation = `-- name: CreateCardConfirmation :one
INSERT INTO card_confirmations (user_id, currency_id, address, verified)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, currency_id, address, verified, image
`

type CreateCardConfirmationParams struct {
	UserID     int64  `json:"user_id"`
	CurrencyID int64  `json:"currency_id"`
	Address    string `json:"address"`
	Verified   bool   `json:"verified"`
}

func (q *Queries) CreateCardConfirmation(ctx context.Context, arg CreateCardConfirmationParams) (CardConfirmation, error) {
	row := q.db.QueryRow(ctx, createCardConfirmation,
		arg.UserID,
		arg.CurrencyID,
		arg.Address,
		arg.Verified,
	)
	var i CardConfirmation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CurrencyID,
		&i.Address,
		&i.Verified,
		&i.Image,
	)
	return i, err
}

const createCurrency = `-- name: CreateCurrency :one
INSERT INTO currencies (code, description)
VALUES ($1, $2)
RETURNING id, code, description
`

type CreateCurrencyParams struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (q *Queries) CreateCurrency(ctx context.Context, arg CreateCurrencyParams) (Currency, error) {
	row := q.db.QueryRow(ctx, createCurrency, arg.Code, arg.Description)
	var i Currency
	err := row.Scan(&i.ID, &i.Code, &i.Description)
	return i, err
}

const createExchanger = `-- name: CreateExchanger :one
INSERT INTO exchangers (
    description,
    inmin,
    require_payment_verification,
    in_currency,
    out_currency
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING id, inmin, description, require_payment_verification, in_currency, out_currency
`

type CreateExchangerParams struct {
	Description                string  `json:"description"`
	Inmin                      float64 `json:"inmin"`
	RequirePaymentVerification bool    `json:"require_payment_verification"`
	InCurrency                 int64   `json:"in_currency"`
	OutCurrency                int64   `json:"out_currency"`
}

func (q *Queries) CreateExchanger(ctx context.Context, arg CreateExchangerParams) (Exchanger, error) {
	row := q.db.QueryRow(ctx, createExchanger,
		arg.Description,
		arg.Inmin,
		arg.RequirePaymentVerification,
		arg.InCurrency,
		arg.OutCurrency,
	)
	var i Exchanger
	err := row.Scan(
		&i.ID,
		&i.Inmin,
		&i.Description,
		&i.RequirePaymentVerification,
		&i.InCurrency,
		&i.OutCurrency,
	)
	return i, err
}

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    operator_id,
    exchanger_id,
    amount_in,
    amount_out,
    cancelled,
    receive_address,
    finished
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, user_id, operator_id, exchanger_id, amount_in, amount_out, receive_address, cancelled, finished
`

type CreateOrderParams struct {
	UserID         int64   `json:"user_id"`
	OperatorID     int64   `json:"operator_id"`
	ExchangerID    int64   `json:"exchanger_id"`
	AmountIn       float64 `json:"amount_in"`
	AmountOut      float64 `json:"amount_out"`
	Cancelled      bool    `json:"cancelled"`
	ReceiveAddress string  `json:"receive_address"`
	Finished       bool    `json:"finished"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder,
		arg.UserID,
		arg.OperatorID,
		arg.ExchangerID,
		arg.AmountIn,
		arg.AmountOut,
		arg.Cancelled,
		arg.ReceiveAddress,
		arg.Finished,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.OperatorID,
		&i.ExchangerID,
		&i.AmountIn,
		&i.AmountOut,
		&i.ReceiveAddress,
		&i.Cancelled,
		&i.Finished,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    email,
    verified,
    passwhash,
    token,
    admin,
    busy
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, email, verified, passwhash, admin, token, busy
`

type CreateUserParams struct {
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
	Passwhash string `json:"passwhash"`
	Token     string `json:"token"`
	Admin     bool   `json:"admin"`
	Busy      bool   `json:"busy"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.Verified,
		arg.Passwhash,
		arg.Token,
		arg.Admin,
		arg.Busy,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Passwhash,
		&i.Admin,
		&i.Token,
		&i.Busy,
	)
	return i, err
}

const getCardConfirmation = `-- name: GetCardConfirmation :one
SELECT id, user_id, currency_id, address, verified, image
FROM card_confirmations
WHERE user_id = $1
  AND currency_id = $2
`

type GetCardConfirmationParams struct {
	UserID     int64 `json:"user_id"`
	CurrencyID int64 `json:"currency_id"`
}

func (q *Queries) GetCardConfirmation(ctx context.Context, arg GetCardConfirmationParams) (CardConfirmation, error) {
	row := q.db.QueryRow(ctx, getCardConfirmation, arg.UserID, arg.CurrencyID)
	var i CardConfirmation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CurrencyID,
		&i.Address,
		&i.Verified,
		&i.Image,
	)
	return i, err
}

const getCurrencyByCode = `-- name: GetCurrencyByCode :one
SELECT id, code, description
FROM currencies
WHERE code = $1
`

func (q *Queries) GetCurrencyByCode(ctx context.Context, code string) (Currency, error) {
	row := q.db.QueryRow(ctx, getCurrencyByCode, code)
	var i Currency
	err := row.Scan(&i.ID, &i.Code, &i.Description)
	return i, err
}

const getCurrencyById = `-- name: GetCurrencyById :one
SELECT id, code, description
FROM currencies
WHERE id = $1
`

func (q *Queries) GetCurrencyById(ctx context.Context, id int64) (Currency, error) {
	row := q.db.QueryRow(ctx, getCurrencyById, id)
	var i Currency
	err := row.Scan(&i.ID, &i.Code, &i.Description)
	return i, err
}

const getExchangerByCurrencyIds = `-- name: GetExchangerByCurrencyIds :one
SELECT id, inmin, description, require_payment_verification, in_currency, out_currency
FROM exchangers
WHERE in_currency = $1
  AND out_currency = $2
`

type GetExchangerByCurrencyIdsParams struct {
	InCurrency  int64 `json:"in_currency"`
	OutCurrency int64 `json:"out_currency"`
}

func (q *Queries) GetExchangerByCurrencyIds(ctx context.Context, arg GetExchangerByCurrencyIdsParams) (Exchanger, error) {
	row := q.db.QueryRow(ctx, getExchangerByCurrencyIds, arg.InCurrency, arg.OutCurrency)
	var i Exchanger
	err := row.Scan(
		&i.ID,
		&i.Inmin,
		&i.Description,
		&i.RequirePaymentVerification,
		&i.InCurrency,
		&i.OutCurrency,
	)
	return i, err
}

const getExchangerById = `-- name: GetExchangerById :one
SELECT id, inmin, description, require_payment_verification, in_currency, out_currency
FROM exchangers
WHERE id = $1
`

func (q *Queries) GetExchangerById(ctx context.Context, id int64) (Exchanger, error) {
	row := q.db.QueryRow(ctx, getExchangerById, id)
	var i Exchanger
	err := row.Scan(
		&i.ID,
		&i.Inmin,
		&i.Description,
		&i.RequirePaymentVerification,
		&i.InCurrency,
		&i.OutCurrency,
	)
	return i, err
}

const getFreeAdmins = `-- name: GetFreeAdmins :many
SELECT id, email, verified, passwhash, admin, token, busy
FROM users
WHERE admin = TRUE
  AND busy = FALSE
`

func (q *Queries) GetFreeAdmins(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getFreeAdmins)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Verified,
			&i.Passwhash,
			&i.Admin,
			&i.Token,
			&i.Busy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrder = `-- name: GetOrder :one
SELECT id, user_id, operator_id, exchanger_id, amount_in, amount_out, receive_address, cancelled, finished
FROM orders
WHERE id = $1
`

func (q *Queries) GetOrder(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRow(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.OperatorID,
		&i.ExchangerID,
		&i.AmountIn,
		&i.AmountOut,
		&i.ReceiveAddress,
		&i.Cancelled,
		&i.Finished,
	)
	return i, err
}

const getOrders = `-- name: GetOrders :many
SELECT id, user_id, operator_id, exchanger_id, amount_in, amount_out, receive_address, cancelled, finished
FROM orders
WHERE finished = false
  AND operator_id = $1
`

func (q *Queries) GetOrders(ctx context.Context, operatorID int64) ([]Order, error) {
	rows, err := q.db.Query(ctx, getOrders, operatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.OperatorID,
			&i.ExchangerID,
			&i.AmountIn,
			&i.AmountOut,
			&i.ReceiveAddress,
			&i.Cancelled,
			&i.Finished,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, email, verified, passwhash, admin, token, busy
FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Passwhash,
		&i.Admin,
		&i.Token,
		&i.Busy,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, verified, passwhash, admin, token, busy
FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Passwhash,
		&i.Admin,
		&i.Token,
		&i.Busy,
	)
	return i, err
}

const getUserByToken = `-- name: GetUserByToken :one
SELECT id, email, verified, passwhash, admin, token, busy
FROM users
WHERE token = $1
`

func (q *Queries) GetUserByToken(ctx context.Context, token string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByToken, token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Passwhash,
		&i.Admin,
		&i.Token,
		&i.Busy,
	)
	return i, err
}

const listBalances = `-- name: ListBalances :many
SELECT id, user_id, currency_id, balance, address
FROM balances
WHERE user_id = $1
`

func (q *Queries) ListBalances(ctx context.Context, userID int64) ([]Balance, error) {
	rows, err := q.db.Query(ctx, listBalances, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Balance
	for rows.Next() {
		var i Balance
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CurrencyID,
			&i.Balance,
			&i.Address,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCurrencies = `-- name: ListCurrencies :many
SELECT id, code, description
FROM currencies
`

func (q *Queries) ListCurrencies(ctx context.Context) ([]Currency, error) {
	rows, err := q.db.Query(ctx, listCurrencies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Currency
	for rows.Next() {
		var i Currency
		if err := rows.Scan(&i.ID, &i.Code, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExchangers = `-- name: ListExchangers :many
SELECT id, inmin, description, require_payment_verification, in_currency, out_currency
FROM exchangers
`

func (q *Queries) ListExchangers(ctx context.Context) ([]Exchanger, error) {
	rows, err := q.db.Query(ctx, listExchangers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exchanger
	for rows.Next() {
		var i Exchanger
		if err := rows.Scan(
			&i.ID,
			&i.Inmin,
			&i.Description,
			&i.RequirePaymentVerification,
			&i.InCurrency,
			&i.OutCurrency,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, verified, passwhash, admin, token, busy
FROM users
ORDER BY email
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Verified,
			&i.Passwhash,
			&i.Admin,
			&i.Token,
			&i.Busy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeCurrency = `-- name: RemoveCurrency :exec
DELETE FROM currencies
WHERE code = $1
`

func (q *Queries) RemoveCurrency(ctx context.Context, code string) error {
	_, err := q.db.Exec(ctx, removeCurrency, code)
	return err
}

const removeExchanger = `-- name: RemoveExchanger :exec
DELETE FROM exchangers
WHERE in_currency = $1
  AND out_currency = $2
`

type RemoveExchangerParams struct {
	InCurrency  int64 `json:"in_currency"`
	OutCurrency int64 `json:"out_currency"`
}

func (q *Queries) RemoveExchanger(ctx context.Context, arg RemoveExchangerParams) error {
	_, err := q.db.Exec(ctx, removeExchanger, arg.InCurrency, arg.OutCurrency)
	return err
}

const updateBalance = `-- name: UpdateBalance :one
UPDATE balances
SET balance = $3
WHERE id = $1
  AND user_id = $2
RETURNING id, user_id, currency_id, balance, address
`

type UpdateBalanceParams struct {
	ID      int64   `json:"id"`
	UserID  int64   `json:"user_id"`
	Balance float64 `json:"balance"`
}

func (q *Queries) UpdateBalance(ctx context.Context, arg UpdateBalanceParams) (Balance, error) {
	row := q.db.QueryRow(ctx, updateBalance, arg.ID, arg.UserID, arg.Balance)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CurrencyID,
		&i.Balance,
		&i.Address,
	)
	return i, err
}

const updateCardConfirmationImage = `-- name: UpdateCardConfirmationImage :one
UPDATE card_confirmations
SET image = $2,
  verified = $3
WHERE id = $1
RETURNING id, user_id, currency_id, address, verified, image
`

type UpdateCardConfirmationImageParams struct {
	ID       int64  `json:"id"`
	Image    []byte `json:"image"`
	Verified bool   `json:"verified"`
}

func (q *Queries) UpdateCardConfirmationImage(ctx context.Context, arg UpdateCardConfirmationImageParams) (CardConfirmation, error) {
	row := q.db.QueryRow(ctx, updateCardConfirmationImage, arg.ID, arg.Image, arg.Verified)
	var i CardConfirmation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CurrencyID,
		&i.Address,
		&i.Verified,
		&i.Image,
	)
	return i, err
}

const updateOrderFinished = `-- name: UpdateOrderFinished :one
UPDATE orders
SET finished = TRUE
WHERE id = $1
RETURNING id, user_id, operator_id, exchanger_id, amount_in, amount_out, receive_address, cancelled, finished
`

func (q *Queries) UpdateOrderFinished(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRow(ctx, updateOrderFinished, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.OperatorID,
		&i.ExchangerID,
		&i.AmountIn,
		&i.AmountOut,
		&i.ReceiveAddress,
		&i.Cancelled,
		&i.Finished,
	)
	return i, err
}

const updateUserBusy = `-- name: UpdateUserBusy :one
UPDATE users
SET busy = $2
WHERE email = $1
RETURNING id, email, verified, passwhash, admin, token, busy
`

type UpdateUserBusyParams struct {
	Email string `json:"email"`
	Busy  bool   `json:"busy"`
}

func (q *Queries) UpdateUserBusy(ctx context.Context, arg UpdateUserBusyParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserBusy, arg.Email, arg.Busy)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Passwhash,
		&i.Admin,
		&i.Token,
		&i.Busy,
	)
	return i, err
}

const updateUserToken = `-- name: UpdateUserToken :one
UPDATE users
SET token = $2
WHERE id = $1
RETURNING id, email, verified, passwhash, admin, token, busy
`

type UpdateUserTokenParams struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

func (q *Queries) UpdateUserToken(ctx context.Context, arg UpdateUserTokenParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserToken, arg.ID, arg.Token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Passwhash,
		&i.Admin,
		&i.Token,
		&i.Busy,
	)
	return i, err
}

const updateUserTokenAndPassHash = `-- name: UpdateUserTokenAndPassHash :one
UPDATE users
SET token = $2,
  passwhash = $3
WHERE email = $1
RETURNING id, email, verified, passwhash, admin, token, busy
`

type UpdateUserTokenAndPassHashParams struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	Passwhash string `json:"passwhash"`
}

func (q *Queries) UpdateUserTokenAndPassHash(ctx context.Context, arg UpdateUserTokenAndPassHashParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserTokenAndPassHash, arg.Email, arg.Token, arg.Passwhash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Passwhash,
		&i.Admin,
		&i.Token,
		&i.Busy,
	)
	return i, err
}

const updateUserVerified = `-- name: UpdateUserVerified :one
UPDATE users
SET verified = $2
WHERE email = $1
RETURNING id, email, verified, passwhash, admin, token, busy
`

type UpdateUserVerifiedParams struct {
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}

func (q *Queries) UpdateUserVerified(ctx context.Context, arg UpdateUserVerifiedParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserVerified, arg.Email, arg.Verified)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.Passwhash,
		&i.Admin,
		&i.Token,
		&i.Busy,
	)
	return i, err
}
