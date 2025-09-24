-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_comments (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    creator_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(creator_id) REFERENCES users_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    content text,
    pictures_attached TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Create a trigger that calls the function before each update
CREATE TRIGGER update_post_comments_updated_at BEFORE UPDATE
ON post_comments
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_comments;
-- +goose StatementEnd
