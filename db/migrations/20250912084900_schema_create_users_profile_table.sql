-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_profile (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES users_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    profile_picture VARCHAR(200),
    avatar VARCHAR(25),
    relation_status VARCHAR(12),
    dob DATE,
    bio VARCHAR(250),
    gender VARCHAR(6),
    family_members TEXT[],
    hobbies TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create a trigger that calls the function before each update
CREATE TRIGGER update_users_profile_updated_at BEFORE UPDATE
ON users_profile
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_profile;
-- +goose StatementEnd
