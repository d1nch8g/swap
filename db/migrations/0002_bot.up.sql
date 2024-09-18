CREATE TABLE bot_messages (
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    message TEXT NOT NULL,
    checked BOOLEAN NOT NULL
);