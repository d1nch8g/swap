CREATE TABLE currencies (
  id BIGSERIAL PRIMARY KEY,
  code TEXT UNIQUE NOT NULL,
  description TEXT UNIQUE NOT NULL
);
CREATE TABLE exchangers (
  id BIGSERIAL PRIMARY KEY,
  inmin DOUBLE PRECISION NOT NULL,
  description TEXT NOT NULL,
  require_payment_verification BOOLEAN NOT NULL,
  input BIGSERIAL NOT NULL,
  CONSTRAINT fk_curr_in FOREIGN KEY(input) REFERENCES currencies(id),
  output BIGSERIAL NOT NULL,
  CONSTRAINT fk_curr_out FOREIGN KEY(output) REFERENCES currencies(id)
);
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  verified BOOLEAN NOT NULL,
  passwhash TEXT NOT NULL,
  admin BOOLEAN NOT NULL,
  token TEXT NOT NULL,
  busy BOOLEAN NOT NULL
);
CREATE TABLE balances (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
  currency_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_currency_id FOREIGN KEY(currency_id) REFERENCES currencies(id),
  balance DOUBLE PRECISION NOT NULL,
  address TEXT UNIQUE NOT NULL
);
CREATE TABLE payment_confirmations (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
  currency_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_currency_id FOREIGN KEY(currency_id) REFERENCES currencies(id),
  address TEXT UNIQUE NOT NULL,
  verified BOOLEAN NOT NULL,
  image BYTEA
);
CREATE TABLE orders (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
  operator_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_operator FOREIGN KEY(operator_id) REFERENCES users(id),
  exchanger_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_exchanger_id FOREIGN KEY(exchanger_id) REFERENCES exchangers(id),
  amount_in DOUBLE PRECISION NOT NULL,
  amount_out DOUBLE PRECISION NOT NULL,
  finished BOOLEAN NOT NULL
);