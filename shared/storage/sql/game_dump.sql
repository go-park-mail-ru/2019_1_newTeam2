-- set global innodb_large_prefix = 'OFF';

DROP DATABASE IF EXISTS game_wordtrainer;
CREATE DATABASE game_wordtrainer DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

GRANT ALL PRIVILEGES ON game_wordtrainer.* TO 'root_use'@'%';
FLUSH PRIVILEGES;

USE game_wordtrainer;

DROP TABLE IF EXISTS russian_english_words;

CREATE TABLE russian_english_words (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    Word VARCHAR(256),
    Translate VARCHAR(256)
);

INSERT INTO russian_english_words (Word, Translate) VALUES ('be', _utf8'быть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('have', _utf8'иметь');
INSERT INTO russian_english_words (Word, Translate) VALUES ('do', _utf8'делать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('say', _utf8'говорить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('go', _utf8'идти');
INSERT INTO russian_english_words (Word, Translate) VALUES ('get', _utf8'получать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('make', _utf8'делать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('know', _utf8'знать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('think', _utf8'думать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('take', _utf8'брать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('see', _utf8'видеть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('come', _utf8'приходить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('want', _utf8'хотеть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('use', _utf8'использовать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('find', _utf8'находить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('give', _utf8'давать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('tell', _utf8'рассказывать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('work', _utf8'работать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('call', _utf8'звать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('try', _utf8'пытаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('ask', _utf8'просить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('need', _utf8'нуждаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('feel', _utf8'чувствовать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('become', _utf8'становиться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('leave', _utf8'оставлять');
INSERT INTO russian_english_words (Word, Translate) VALUES ('put', _utf8'класть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('mean', _utf8'значить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('keep', _utf8'хранить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('let', _utf8'позволять');
INSERT INTO russian_english_words (Word, Translate) VALUES ('begin', _utf8'начинать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('seem', _utf8'казаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('help', _utf8'помогать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('show', _utf8'показывать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('hear', _utf8'слышать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('play', _utf8'играть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('run', _utf8'бежать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('move', _utf8'двигаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('live', _utf8'жить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('believe', _utf8'верить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('bring', _utf8'приносить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('happen', _utf8'случаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('write', _utf8'писать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('sit', _utf8'сидеть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('stand', _utf8'стоять');
INSERT INTO russian_english_words (Word, Translate) VALUES ('lose', _utf8'терять');
INSERT INTO russian_english_words (Word, Translate) VALUES ('pay', _utf8'платить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('meet', _utf8'встречать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('include', _utf8'включать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('continue', _utf8'продолжать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('set', _utf8'устанавливать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('learn', _utf8'учить');
