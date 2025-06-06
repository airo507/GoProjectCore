-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS users_email_idx ON users(email);
CREATE INDEX IF NOT EXISTS users_created_at_idx ON users(created_at);

CREATE TABLE IF NOT EXISTS posts
(
    id SERIAL PRIMARY KEY,
    author_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    likes INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users (id)
);
CREATE INDEX IF NOT EXISTS posts_created_at_idx ON posts(created_at);
CREATE INDEX IF NOT EXISTS posts_author_created_idx ON posts(author_id, created_at);

CREATE TABLE IF NOT EXISTS comments
(
    id SERIAL PRIMARY KEY,
    author_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users (id),
    FOREIGN KEY (post_id) REFERENCES posts (id)
);
CREATE INDEX IF NOT EXISTS comments_created_at_idx ON comments(created_at);
CREATE INDEX IF NOT EXISTS comments_post_created_idx ON comments(post_id, created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE posts;
DROP TABLE comments;
-- +goose StatementEnd
