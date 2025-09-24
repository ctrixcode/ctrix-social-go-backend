-- +goose Up
-- +goose StatementBegin
INSERT INTO users_profile (id, first_name, last_name, avatar, relation_status, dob, bio, gender, family_members, hobbies) VALUES
('1', 'Anmol', 'Jain', 'avatar_aj', 'Single', '1990-01-01', 'Passionate developer and tech enthusiast.', 'Male', '{}', '{coding,gaming,reading}'),
('2', 'Ctrix', 'User', 'avatar_ctrix', 'Married', '1985-05-15', 'Building awesome things with Go.', 'Female', '{}', '{traveling,photography,cooking}'),
('3', 'Mike', 'Jones', 'avatar_mj', 'Single', '1992-03-20', 'Loves hiking and photography.', 'Male', '{}', '{hiking,photography,climbing}'),
('4', 'Sara', 'Connor', 'avatar_sc', 'In a relationship', '1988-11-10', 'Future of humanity depends on me.', 'Female', '{}', '{survival,tactics,running}'),
('5', 'Chris', 'Evans', 'avatar_ce', 'Single', '1981-06-13', 'Avenger by day, actor by night.', 'Male', '{}', '{acting,fitness,charity}'),
('6', 'Lisa', 'Brown', 'avatar_lb', 'Married', '1995-08-22', 'Enjoying life one day at a time.', 'Female', '{}', '{gardening,baking,yoga}'),
('7', 'David', 'Lee', 'avatar_dl', 'Single', '1993-04-05', 'Exploring new technologies.', 'Male', '{}', '{tech,gadgets,cycling}'),
('8', 'Emily', 'White', 'avatar_ew', 'In a relationship', '1991-09-30', 'Artist and nature lover.', 'Female', '{}', '{painting,nature,meditation}'),
('9', 'Peter', 'Parker', 'avatar_pp', 'Single', '2000-08-01', 'Your friendly neighborhood Spider-Man.', 'Male', '{}', '{web-slinging,science,photography}'),
('10', 'Mary', 'Jane', 'avatar_mjn', 'In a relationship', '2001-03-17', 'Aspiring actress and model.', 'Female', '{}', '{acting,modeling,dancing}');
-- +goose StatementEnd