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
  in_currency BIGSERIAL NOT NULL,
  CONSTRAINT fk_curr_in FOREIGN KEY(in_currency) REFERENCES currencies(id),
  out_currency BIGSERIAL NOT NULL,
  CONSTRAINT fk_curr_out FOREIGN KEY(out_currency) REFERENCES currencies(id),
  unique(in_currency, out_currency)
);
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  verified BOOLEAN NOT NULL,
  passwhash TEXT NOT NULL,
  admin BOOLEAN NOT NULL,
  operator BOOLEAN NOT NULL,
  token TEXT NOT NULL,
  busy BOOLEAN NOT NULL
);
CREATE TABLE balances (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
  currency_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_currency_id FOREIGN KEY(currency_id) REFERENCES currencies(id),
  unique (user_id, currency_id),
  balance DOUBLE PRECISION NOT NULL,
  address TEXT UNIQUE NOT NULL
);
CREATE TABLE card_confirmations (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
  currency_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_currency_id FOREIGN KEY(currency_id) REFERENCES currencies(id),
  unique (user_id, currency_id),
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
  receive_address TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  cancelled BOOLEAN NOT NULL,
  finished BOOLEAN NOT NULL,
  confirm_image BYTEA,
  payment_confirmed BOOLEAN NOT NULL
);