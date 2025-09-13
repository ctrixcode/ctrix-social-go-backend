-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_auth (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(50) UNIQUE NOT NULL,
    username VARCHAR(30) UNIQUE NOT NULL,
    password VARCHAR(75) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- Create a trigger that calls the function before each update
CREATE TRIGGER update_users_auth_updated_at BEFORE UPDATE
ON users_auth
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_auth;
-- +goose StatementEnd