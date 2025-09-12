-- +goose Up
-- +goose StatementBegin
INSERT INTO user_profile (id, first_name, last_name, profile_picture, avatar, last_seen, post_count, followers, followings) VALUES
('1', 'Anmol', 'Jain', 'https://example.com/profile/aj_noob.jpg', 'avatar_aj', NOW(), 5, '{}', '{}'),
('2', 'Ctrix', 'User', 'https://example.com/profile/ctrix.jpg', 'avatar_ctrix', NOW(), 10, '{}', '{}'),
('3', 'Mike', 'Jones', 'https://example.com/profile/mike.jpg', 'avatar_mj', NOW(), 2, '{}', '{}'),
('4', 'Sara', 'Connor', 'https://example.com/profile/sara.jpg', 'avatar_sc', NOW(), 8, '{}', '{}'),
('5', 'Chris', 'Evans', 'https://example.com/profile/chris.jpg', 'avatar_ce', NOW(), 3, '{}', '{}'),
('6', 'Lisa', 'Brown', 'https://example.com/profile/lisa.jpg', 'avatar_lb', NOW(), 7, '{}', '{}'),
('7', 'David', 'Lee', 'https://example.com/profile/david.jpg', 'avatar_dl', NOW(), 1, '{}', '{}'),
('8', 'Emily', 'White', 'https://example.com/profile/emily.jpg', 'avatar_ew', NOW(), 6, '{}', '{}'),
('9', 'Peter', 'Parker', 'https://example.com/profile/peter.jpg', 'avatar_pp', NOW(), 4, '{}', '{}'),
('10', 'Mary', 'Jane', 'https://example.com/profile/mary.jpg', 'avatar_mjn', NOW(), 9, '{}', '{}');
-- +goose StatementEnd