USE game_wordtrainer;

DROP TABLE IF EXISTS russian_english_words;

CREATE TABLE russian_english_words (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    Word VARCHAR(256),
    Translate VARCHAR(256)
);