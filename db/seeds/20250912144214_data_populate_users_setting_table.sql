-- +goose Up
-- +goose StatementBegin
INSERT INTO users_setting (id, block_user, hide_post, hide_story, show_online) VALUES
('1', '{"2"}', '{}', '{}', true),
('2', '{"4"}', '{}', '{}', true),
('3', '{}', '{}', '{}', false),
('4', '{"1", "3"}', '{}', '{}', true);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users_setting where true;
-- +goose StatementEnd
