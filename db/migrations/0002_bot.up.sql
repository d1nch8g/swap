CREATE TABLE bot_messages (
    user_id BIGSERIAL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    order_id BIGSERIAL,
    CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(id),
    message TEXT NOT NULL,
    checked BOOLEAN NOT NULL
);