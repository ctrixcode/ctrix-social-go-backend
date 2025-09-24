-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_likes (
    user_id VARCHAR(50),
    FOREIGN KEY(user_id) REFERENCES users_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    post_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (post_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_likes;
-- +goose StatementEnd
