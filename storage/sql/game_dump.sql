-- set global innodb_large_prefix = 'OFF';

DROP DATABASE IF EXISTS game_wordtrainer;
CREATE DATABASE game_wordtrainer;

ALTER DATABASE game_wordtrainer CHARACTER SET utf8 COLLATE utf8_general_ci;

GRANT ALL PRIVILEGES ON game_wordtrainer.* TO 'root_use'@'%';
FLUSH PRIVILEGES;

USE game_wordtrainer;

DROP TABLE IF EXISTS russian_english_words;

CREATE TABLE russian_english_words (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    Word VARCHAR(256),
    Translate VARCHAR(256)
);

INSERT INTO russian_english_words (Word, Translate) VALUES ('be', 'быть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('have', 'иметь');
INSERT INTO russian_english_words (Word, Translate) VALUES ('do', 'делать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('say', 'говорить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('go', 'идти');
INSERT INTO russian_english_words (Word, Translate) VALUES ('get', 'получать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('make', 'делать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('know', 'знать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('think', 'думать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('take', 'брать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('see', 'видеть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('come', 'приходить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('want', 'хотеть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('use', 'использовать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('find', 'находить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('give', 'давать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('tell', 'рассказывать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('work', 'работать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('call', 'звать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('try', 'пытаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('ask', 'просить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('need', 'нуждаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('feel', 'чувствовать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('become', 'становиться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('leave', 'оставлять');
INSERT INTO russian_english_words (Word, Translate) VALUES ('put', 'класть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('mean', 'значить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('keep', 'хранить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('let', 'позволять');
INSERT INTO russian_english_words (Word, Translate) VALUES ('begin', 'начинать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('seem', 'казаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('help', 'помогать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('show', 'показывать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('hear', 'слышать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('play', 'играть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('run', 'бежать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('move', 'двигаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('live', 'жить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('believe', 'верить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('bring', 'приносить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('happen', 'случаться');
INSERT INTO russian_english_words (Word, Translate) VALUES ('write', 'писать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('sit', 'сидеть');
INSERT INTO russian_english_words (Word, Translate) VALUES ('stand', 'стоять');
INSERT INTO russian_english_words (Word, Translate) VALUES ('lose', 'терять');
INSERT INTO russian_english_words (Word, Translate) VALUES ('pay', 'платить');
INSERT INTO russian_english_words (Word, Translate) VALUES ('meet', 'встречать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('include', 'включать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('continue', 'продолжать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('set', 'устанавливать');
INSERT INTO russian_english_words (Word, Translate) VALUES ('learn', 'учить');
