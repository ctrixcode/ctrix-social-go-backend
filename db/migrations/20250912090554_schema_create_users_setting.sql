-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_setting (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES users_auth ON DELETE CASCADE ON UPDATE CASCADE,
    block_user TEXT[],
    hide_post TEXT[],
    hide_story TEXT[],
    show_online BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create a trigger that calls the function before each update
CREATE TRIGGER update_users_setting BEFORE UPDATE
ON users_setting
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_setting;
-- +goose StatementEnd
