CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR (50) UNIQUE NOT NULL,
  card VARCHAR (50) UNIQUE NOT NULL,
  verified BOOLEAN NOT NULL
);
CREATE TABLE admins (
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR (50) UNIQUE NOT NULL,
  PASSWHASH VARCHAR (50) NOT NULL
);
CREATE TABLE orders (
  id BIGSERIAL PRIMARY KEY,
  give TEXT NOT NULL,
  receive TEXT NOT NULL,
  user_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);
CREATE TABLE exchangers (
  id BIGSERIAL PRIMARY KEY,
  inmin DECIMAL NOT NULL,
  inmax DECIMAL NOT NULL,
  reserve DECIMAL NOT NULL,
  rate DECIMAL NOT NULL,
  change TEXT NOT NULL
);
CREATE TABLE order_chats (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
  order_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(id)
);