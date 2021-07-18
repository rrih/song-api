CREATE DATABASE IF NOT EXISTS managedby_db;
USE managedby_db;

CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(64),
    email VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    created DATETIME NOT NULL,
    modified DATETIME NOT NULL,
    PRIMARY KEY (id)
);

-- TODO: 楽曲用テーブル作成
-- CREATE TABLE IF NOT EXISTS musics (
--     id INT NOT NULL AUTO_INCREMENT,
--     name VARCHAR(100),

-- )

INSERT INTO users (id, name, email, password, created, modified) VALUES (1, 'admin', 'rsklv@test.com', 'testtest', '2020-01-01 00:00:00', '2020-01-01 00:00:00');
INSERT INTO users (id, name, email, password, created, modified) VALUES (2, 'origa', 'test@gmail.com', 'testtest', '2020-01-01 00:00:00', '2020-01-01 00:00:00');
INSERT INTO users (id, name, email, password, created, modified) VALUES (3, 'ryohei', 'k.ryohei@test.co.jp', 'testtest', '2020-02-01 00:00:00', '2020-02-01 00:00:00');
