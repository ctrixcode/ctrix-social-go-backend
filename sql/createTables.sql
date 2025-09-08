-- Enable the uuid-ossp extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_auth (
    id VARCHAR(50) PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    username VARCHAR(30) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_profile (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    profile_picture VARCHAR(200),
    avatar VARCHAR(25),
    last_seen TIMESTAMP,
    post_count INT,
    followers TEXT[],
    followings TEXT[]
);


CREATE TABLE IF NOT EXISTS user_additional_info (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES user_auth ON DELETE CASCADE ON UPDATE CASCADE,
    hobbies TEXT[],
    family_members TEXT[],
    relation_status VARCHAR(12),
    dob DATE,
    bio VARCHAR(250),
    gender VARCHAR(6)
);

CREATE TABLE IF NOT EXISTS user_settings (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES user_auth ON DELETE CASCADE ON UPDATE CASCADE,
    hide_post TEXT[],
    hide_story TEXT[],
    block_user TEXT[],
    show_online BOOLEAN
);

CREATE TABLE IF NOT EXISTS user_data (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES user_auth ON DELETE CASCADE ON UPDATE CASCADE,
    posts TEXT[],
    stories TEXT[],
    notes TEXT[]
);

CREATE TABLE IF NOT EXISTS user_posts (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    creator_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    group_id VARCHAR(50),
    text_content TEXT,
    pictures_attached TEXT[],
    comments TEXT[]
);

CREATE TABLE IF NOT EXISTS user_post_like (
    user_id VARCHAR(50),
    FOREIGN KEY(user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    post_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(post_id) REFERENCES user_posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    like_type VARCHAR(20),
    UNIQUE (post_id, user_id)
);

CREATE TABLE IF NOT EXISTS user_post_comments (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(post_id) REFERENCES user_posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    creator_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    content text,
    pictures_attached TEXT[],
    nested_comments TEXT[]
);

CREATE TABLE IF NOT EXISTS user_comment_like (
    user_id VARCHAR(50),
    FOREIGN KEY(user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    comment_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(comment_id) REFERENCES user_post_comments(id) ON DELETE CASCADE ON UPDATE CASCADE,
    like_type VARCHAR(20),
    UNIQUE (comment_id, user_id)
);
