-- +goose Up
CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    author INT NOT NULL,
    label TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    chat_id INT NOT NULL,
    author INT NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE chats;
DROP TABLE messages;
