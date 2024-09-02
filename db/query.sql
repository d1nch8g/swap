-- name: CreateCurrency :one
INSERT INTO currencies (
    code,
    description,
    bestchange_id,
    accepted_window,
    require_payment_verification
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: ListCurrencies :many
SELECT *
FROM currencies;
-- name: GetCurrencyByCode :many
SELECT *
FROM currencies
WHERE code = $1;
-- name: RemoveCurrency :exec
DELETE FROM currencies
WHERE code = $1;
-- name: CreateExchanger :one
INSERT INTO exchangers (
    rate,
    description,
    inmin,
    input,
    output
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: UpdateExchangerRate :one
UPDATE exchangers
set rate = $2
WHERE id = $1
RETURNING *;
-- name: RemoveExchanger :exec
DELETE FROM exchangers
WHERE input = $1
  AND output = $2;
-- name: ListExchangers :many
SELECT *
FROM exchangers;
-- name: CreateUser :one
INSERT INTO users (
    email,
    verified,
    passwhash,
    token,
    admin,
    busy
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: UpdateUserBusy :one
UPDATE users
SET busy = $2
WHERE email = $1
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;
-- name: GetUserByToken :one
SELECT *
FROM users
WHERE token = $1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY email;
-- name: UpdateUserVerified :one
UPDATE users
SET verified = $2
WHERE email = $1
RETURNING *;
-- name: UpdateUserToken :one
UPDATE users
SET token = $2
WHERE id = $1
RETURNING *;
-- name: CreateBalance :one
INSERT INTO balances (user_id, currency_id, balance, address)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateBalance :one
UPDATE balances
SET balance = $3
WHERE id = $1
  AND user_id = $2
RETURNING *;
-- name: ListBalances :many
SELECT *
FROM balances
WHERE user_id = $1;
-- name: CreatePaymentConfirmation :one
INSERT INTO payment_confirmations (user_id, currency_id, address, verified)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetPaymentConfirmation :one
SELECT *
FROM payment_confirmations
WHERE id = $1;
-- name: UpdatePaymentConfirmationImage :one
UPDATE payment_confirmations
SET image = $2,
  verified = $3
WHERE id = $1
RETURNING *;
-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    operator_id,
    exchanger_id,
    amount_in,
    amount_out,
    finished
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: OrdersUnfinished :many
SELECT *
FROM orders
WHERE finished = false;