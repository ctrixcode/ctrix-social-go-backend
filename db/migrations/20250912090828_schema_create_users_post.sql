-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_post (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    creator_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES users_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    group_id VARCHAR(50),
    text_content TEXT,
    pictures_attached TEXT[],
    comments TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Create a trigger that calls the function before each update
CREATE TRIGGER update_users_post_updated_at BEFORE UPDATE
ON users_post
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_post;
-- +goose StatementEnd
