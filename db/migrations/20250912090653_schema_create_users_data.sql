-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_data (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES users_auth ON DELETE CASCADE ON UPDATE CASCADE,
    posts TEXT[],
    stories TEXT[],
    notes TEXT[],
    last_seen TIMESTAMP,
    followers TEXT[],
    followings TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create a trigger that calls the function before each update
CREATE TRIGGER update_users_data_updated_at BEFORE UPDATE
ON users_data
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_data;
-- +goose StatementEnd
