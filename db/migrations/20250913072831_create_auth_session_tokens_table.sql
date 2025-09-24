-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth_session_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id VARCHAR(50) NOT NULL,
    jti TEXT UNIQUE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    is_used BOOLEAN DEFAULT FALSE NOT NULL,
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users_auth(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Create a trigger that calls the function before each update
CREATE TRIGGER update_auth_session_tokens_updated_at BEFORE UPDATE
ON auth_session_tokens
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth_session_tokens;
-- +goose StatementEnd