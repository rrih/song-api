CREATE DATABASE IF NOT EXISTS managedby_db;
USE managedby_db;

-- ユーザーの情報
CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT UNIQUE,
    name VARCHAR(64) UNIQUE,
    email VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(64) NOT NULL,
    is_admin TINYINT(1) NOT NULL DEFAULT 0, -- admin であるか
    deleted DATETIME DEFAULT NULL,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 楽曲の情報
CREATE TABLE IF NOT EXISTS music (
    id INT NOT NULL AUTO_INCREMENT UNIQUE,
    registered_user_id INT NOT NULL,
    category_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    singer_name VARCHAR(100) NOT NULL, -- 歌手名
    composer_name VARCHAR(100) NOT NULL, -- 作曲者名
    source VARCHAR(100) NOT NULL, -- 楽曲を使用したコンテンツ名
    url VARCHAR(256), -- 参考 URL
    is_anime_video_dam TINYINT(1) NOT NULL DEFAULT 0, -- LIVE DAM の場合アニメ映像が存在するか
    is_anime_video_joy TINYINT(1) NOT NULL DEFAULT 0, -- JOYSOUND の場合アニメ映像が存在するか
    is_official_video_dam TINYINT(1) NOT NULL DEFAULT 0,
    is_official_video_joy TINYINT(1) NOT NULL DEFAULT 0,
    start_singing VARCHAR(200), -- 歌い出し歌詞
    deleted DATETIME DEFAULT NULL,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 楽曲のカラオケスコアの情報
CREATE TABLE IF NOT EXISTS score (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    another_user_id INT, -- デュエットの場合
    music_id INT NOT NULL,
    vibrato int NOT NULL DEFAULT 0, -- ビブラート
    kobushi int NOT NULL DEFAULT 0, -- こぶし
    fall int NOT NULL DEFAULT 0, -- フォール
    shakuri int NOT NULL DEFAULT 0, -- しゃくり
    is_dam TINYINT(1) NOT NULL DEFAULT 0, -- LIVE DAM であるか
    is_joy TINYINT(1) NOT NULL DEFAULT 0, -- JOYSOUND であるか
    is_able_to_song_by_man TINYINT(1) NOT NULL DEFAULT 0, -- 男が歌える音程か
    is_able_to_song_by_woman TINYINT(1) NOT NULL DEFAULT 0, -- 女が歌える音程か
    is_fav TINYINT(1) NOT NULL DEFAULT 0, -- お気に入りかどうか
    key_number INT, -- キー設定
    score FLOAT NOT NULL,
    text VARCHAR(1000),
    deleted DATETIME DEFAULT NULL,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- スコアとコメントの中間テーブル
CREATE TABLE IF NOT EXISTS score_comments (
    id INT NOT NULL AUTO_INCREMENT UNIQUE,
    score_id INT NOT NULL,
    comment_id INT NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- コメントの情報
CREATE TABLE IF NOT EXISTS comments (
    id INT NOT NULL AUTO_INCREMENT UNIQUE,
    text VARCHAR(100) NOT NULL,
    deleted DATETIME DEFAULT NULL,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- カテゴリーの情報
CREATE TABLE IF NOT EXISTS categories (
    id INT NOT NULL AUTO_INCREMENT,
    parent_id INT,
    name VARCHAR(1000) NOT NULL,
    deleted DATETIME DEFAULT NULL,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO users (id, name, email, password, is_admin, deleted, created, modified) VALUES (1, 'admin', 'rsklv@test.com', 'testtest', 0, null, '2020-01-01 00:00:00', '2020-01-01 00:00:00');
INSERT INTO users (id, name, email, password, is_admin, deleted, created, modified) VALUES (2, 'origa', 'test@gmail.com', 'testtest', 0, null, '2020-01-01 00:00:00', '2020-01-01 00:00:00');
INSERT INTO users (id, name, email, password, is_admin, deleted, created, modified) VALUES (3, 'ryohei2', 'k.ryohei@test.co.jp', 'testtest', 0, null, '2020-02-01 00:00:00', '2020-02-01 00:00:00');

INSERT INTO music (
    id,
    registered_user_id,
    category_id,
    name,
    singer_name,
    composer_name,
    source,
    url,
    is_anime_video_dam,
    is_anime_video_joy,
    is_official_video_dam,
    is_official_video_joy,
    start_singing,
    deleted,
    created,
    modified
) VALUES (
    1,
    1,
    1,
    '三原色',
    'YOASOBI',
    'YOASOBI',
    'ahamoのCM',
    'https://www.youtube.com/watch?v=nhOhFOoURnE',
    0,
    0,
    0,
    0,
    'どこかで途切れた物語',
    null,
    '2021-07-19 00:00:00',
    '2021-07-19 00:00:00'
);

INSERT INTO score (
    id, user_id, another_user_id, music_id, vibrato, kobushi, fall, shakuri, is_dam, is_joy, is_able_to_song_by_man, is_able_to_song_by_woman, is_fav, key_number, score, text, deleted, created, modified
) VALUES (
    1, 1, null, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 88.234, '長いフレーズに続くビブラートの耳触りがとても心地よいです。テクニックが際立っていますね。', null, '2021-07-19 00:00:00', '2021-07-19 00:00:00'
);

INSERT INTO score_comments (id, score_id, comment_id) VALUES (1, 1, 1);

INSERT INTO comments (id, text, deleted, created, modified) VALUES (1, '後半、息を多く吸う必要あり。ぎりぎり最高音程出せたけどミックスボイスを使えないとダメな感じ。', null, '2021-07-19 00:00:00', '2021-07-19 00:00:00');

INSERT INTO categories (id, parent_id, name, deleted, created, modified) VALUES (1, null, 'CM楽曲', null, '2021-07-19 00:00:00', '2021-07-19 00:00:00');
