-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_comment_likes (
    user_id VARCHAR(50),
    FOREIGN KEY(user_id) REFERENCES users_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    comment_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(comment_id) REFERENCES post_comments(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (comment_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_comment_likes;
-- +goose StatementEnd
