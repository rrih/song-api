CREATE DATABASE IF NOT EXISTS managedby_db;
USE managedby_db;

-- ユーザーの情報
CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(64),
    email VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    is_admin TINYINT(1) NOT NULL DEFAULT 0, -- admin であるか
    deleted DATETIME,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
);

-- 楽曲の情報
CREATE TABLE IF NOT EXISTS music (
    id INT NOT NULL AUTO_INCREMENT,
    registered_user_id INT NOT NULL,
    name VARCHAR(100),
    is_anime_video_dam TINYINT(1) NOT NULL DEFAULT 0, -- LIVE DAM の場合アニメ映像が存在するか
    is_anime_video_joy TINYINT(1) NOT NULL DEFAULT 0, -- JOYSOUND の場合アニメ映像が存在するか
    is_official_video_dam TINYINT(1) NOT NULL DEFAULT 0,
    is_official_video_joy TINYINT(1) NOT NULL DEFAULT 0,
    is_fav TINYINT(1) NOT NULL DEFAULT 0,
    start_singing VARCHAR(200),
    deleted DATETIME,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
);

-- 楽曲のカラオケスコアの情報
CREATE TABLE IF NOT EXISTS score (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    another_user_id INT,
    music_id INT NOT NULL,
    is_dam TINYINT(1) NOT NULL DEFAULT 0, -- LIVE DAM であるか
    is_joy TINYINT(1) NOT NULL DEFAULT 0, -- JOYSOUND であるか
    is_able_to_song_by_man TINYINT(1) NOT NULL DEFAULT 0, -- 男が歌える音程か
    is_able_to_song_by_woman TINYINT(1) NOT NULL DEFAULT 0, -- 女が歌える音程か
    key_number INT, -- キー設定
    score FLOAT NOT NULL,
    key_number INT,
    text VARCHAR(1000),
    deleted DATETIME,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
)

INSERT INTO users (id, name, email, password, is_admin, deleted, created, modified) VALUES (1, 'admin', 'rsklv@test.com', 'testtest', 0, null, '2020-01-01 00:00:00', '2020-01-01 00:00:00');
INSERT INTO users (id, name, email, password, is_admin, deleted, created, modified) VALUES (2, 'origa', 'test@gmail.com', 'testtest', 0, null, '2020-01-01 00:00:00', '2020-01-01 00:00:00');
INSERT INTO users (id, name, email, password, is_admin, deleted, created, modified) VALUES (3, 'ryohei', 'k.ryohei@test.co.jp', 'testtest', 0, null, '2020-02-01 00:00:00', '2020-02-01 00:00:00');
