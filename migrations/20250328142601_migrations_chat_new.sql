-- +goose Up
CREATE TABLE chat_users (
    id SERIAL PRIMARY KEY,
    chat_id INT NOT NULL,
    user_id INT NOT NULL
);

-- +goose Down
DROP TABLE chat_users;
