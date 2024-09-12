-- name: CreateCurrency :one
INSERT INTO currencies (code, description)
VALUES ($1, $2)
RETURNING *;
-- name: ListCurrencies :many
SELECT *
FROM currencies;
-- name: GetCurrencyByCode :one
SELECT *
FROM currencies
WHERE code = $1;
-- name: RemoveCurrency :exec
DELETE FROM currencies
WHERE code = $1;
-- name: GetCurrencyById :one
SELECT *
FROM currencies
WHERE id = $1;
-- name: CreateExchanger :one
INSERT INTO exchangers (
    description,
    inmin,
    require_payment_verification,
    in_currency,
    out_currency
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: RemoveExchanger :exec
DELETE FROM exchangers
WHERE in_currency = $1
  AND out_currency = $2;
-- name: ListExchangers :many
SELECT *
FROM exchangers;
-- name: GetExchangerByCurrencyIds :one
SELECT *
FROM exchangers
WHERE in_currency = $1
  AND out_currency = $2;
-- name: GetExchangerById :one
SELECT *
FROM exchangers
WHERE id = $1;
-- name: CreateUser :one
INSERT INTO users (
    email,
    verified,
    passwhash,
    admin,
    operator,
    token,
    busy
  )
VALUES ($1, $2, $3, $4, $5, $6, $7)
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
-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;
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
-- name: UpdateUserTokenAndPassHash :one
UPDATE users
SET token = $2,
  passwhash = $3
WHERE email = $1
RETURNING *;
-- name: GetFreeAdmins :many
SELECT *
FROM users
WHERE admin = TRUE
  AND busy = FALSE;
-- name: CreateBalance :one
INSERT INTO balances (user_id, currency_id, balance, address)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateBalance :one
UPDATE balances
SET balance = $3,
  address = $4
WHERE id = $1
  AND user_id = $2
RETURNING *;
-- name: ListBalances :many
SELECT *
FROM balances
WHERE user_id = $1;
-- name: GetBalanceById :one
SELECT *
FROM balances
WHERE id = $1
  AND user_id = $2;
-- name: RemoveBalance :exec
DELETE FROM balances
WHERE id = $1
  AND user_id = $2;
-- name: CreateCardConfirmation :one
INSERT INTO card_confirmations (user_id, currency_id, address, verified)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetCardConfirmation :one
SELECT *
FROM card_confirmations
WHERE user_id = $1
  AND currency_id = $2;
-- name: GetCardConfirmations :many
SELECT *
FROM card_confirmations;
-- name: UpdateCardConfirmationImage :one
UPDATE card_confirmations
SET image = $2
WHERE id = $1
RETURNING *;
-- name: UpdateCardConfirmationVerified :one
UPDATE card_confirmations
SET verified = $2
WHERE id = $1
RETURNING *;
-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    operator_id,
    exchanger_id,
    amount_in,
    amount_out,
    cancelled,
    receive_address,
    finished,
    confirm_image,
    payment_confirmed
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;
-- name: GetOrders :many
SELECT *
FROM orders
WHERE finished = false
  AND operator_id = $1;
-- name: GetOrder :one
SELECT *
FROM orders
WHERE id = $1;
-- name: UpdateOrderFinished :one
UPDATE orders
SET finished = TRUE
WHERE id = $1
RETURNING *;
-- name: UpdateOrderCancelled :one
UPDATE orders
SET finished = TRUE,
  cancelled = TRUE
WHERE id = $1
RETURNING *;
-- name: GetOrdersForUser :many
SELECT *
FROM orders
WHERE user_id = $1;
-- name: UpdateOrderPaymentConfirmed :one
UPDATE orders
SET payment_confirmed = $2,
  confirm_image = $3
WHERE id = $1
RETURNING *;