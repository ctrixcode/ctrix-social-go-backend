-- +goose Up
-- +goose StatementBegin
INSERT INTO users_auth (id, email, username, password, created_at) VALUES
('1', 'anmol@anmol.pro', 'aj_noob', 'password123', NOW()),
('2', 'ctrix@ctrix.pro', 'ctrix', 'password123', NOW()),
('3', 'mike.jones@example.com', 'mike.jones', 'password123', NOW()),
('4', 'sara.connor@example.com', 'sara.connor', 'password123', NOW()),
('5', 'chris.evans@example.com', 'chris.evans', 'password123', NOW()),
('6', 'lisa.brown@example.com', 'lisa.brown', 'password123', NOW()),
('7', 'david.lee@example.com', 'david.lee', 'password123', NOW()),
('8', 'emily.white@example.com', 'emily.white', 'password123', NOW()),
('9', 'peter.parker@example.com', 'peter.parker', 'password123', NOW()),
('10', 'mary.jane@example.com', 'mary.jane', 'password123', NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users_auth where true;
-- +goose StatementEnd