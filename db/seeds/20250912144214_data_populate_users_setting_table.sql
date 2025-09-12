-- +goose Up
-- +goose StatementBegin
INSERT INTO users_setting (id, block_user, hide_post, hide_story, show_online, created_at) VALUES
('1', '{"2"}', '{}', '{}', true, NOW()),
('2', '{"4"}', '{}', '{}', true, NOW()),
('3', '{}', '{}', '{}', false, NOW()),
('4', '{"1", "3"}', '{}', '{}', true, NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users_setting where true;
-- +goose StatementEnd