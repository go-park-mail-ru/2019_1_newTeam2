-- set global innodb_large_prefix = 'OFF';

DROP DATABASE IF EXISTS game_wordtrainer;
CREATE DATABASE game_wordtrainer;

GRANT ALL PRIVILEGES ON game_wordtrainer.* TO 'root_use'@'%';
FLUSH PRIVILEGES;

USE game_wordtrainer;

DROP TABLE IF EXISTS russian_english_words;

CREATE TABLE russian_english_words (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    Word VARCHAR(256),
    Translate VARCHAR(256)
);