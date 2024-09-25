CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    uuid TEXT NOT NULL,
    resolved BOOLEAN NOT NULL
);
CREATE TABLE chat_messages (
    chat_id BIGSERIAL NOT NULL,
    CONSTRAINT fk_chat FOREIGN KEY(chat_id) REFERENCES chats(id),
    outgoing BOOLEAN NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);