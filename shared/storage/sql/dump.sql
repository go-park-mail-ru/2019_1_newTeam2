-- set global innodb_large_prefix = 'OFF';

DROP DATABASE IF EXISTS wordtrainer;
CREATE DATABASE wordtrainer DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

GRANT ALL PRIVILEGES ON wordtrainer.* TO 'root_use'@'%';
FLUSH PRIVILEGES;

USE wordtrainer;

DROP TABLE IF EXISTS dictionary_to_library;
DROP TABLE IF EXISTS cards_library;
DROP TABLE IF EXISTS dictionary;
DROP TABLE IF EXISTS card;
DROP TABLE IF EXISTS word;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS language;

CREATE TABLE language (
	ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(128) NOT NULL UNIQUE
);

INSERT INTO language (name) VALUES ("English");
INSERT INTO language (name) VALUES ("Russian");

CREATE TABLE user (
	ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	Username VARCHAR(128) NOT NULL UNIQUE,
	Email VARCHAR(128) NOT NULL,
	Password VARCHAR(128) NOT NULL,
	LangID INT UNSIGNED NOT NULL,
	PronounceON TINYINT NOT NULL,
	Score INT UNSIGNED NOT NULL,
	AvatarPath VARCHAR(500) NOT NULL,
	FOREIGN KEY (LangID) REFERENCES language (ID)
);

CREATE TABLE dictionary (
	ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(128) NOT NULL,
	description TEXT,
	UserID INT UNSIGNED NOT NULL,
	FOREIGN KEY (UserID) REFERENCES user (ID)
);

CREATE TABLE word (
	ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(128) NOT NULL,
	LangID INT UNSIGNED NOT NULL,
	FOREIGN KEY (LangID) REFERENCES language (ID)
);

CREATE TABLE card (
	ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	word INT UNSIGNED NOT NULL,
	translation INT UNSIGNED NOT NULL,
	FOREIGN KEY (word) REFERENCES word (ID),
	FOREIGN KEY (translation) REFERENCES word (ID)
);

CREATE TABLE cards_library (
	ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	card_id INT UNSIGNED NOT NULL,
	count INT UNSIGNED NOT NULL,
	guessed INT UNSIGNED NOT NULL DEFAULT 0,
	seen INT UNSIGNED NOT NULL DEFAULT 1,
	if_seen BOOL NOT NULL DEFAULT FALSE,
	FOREIGN KEY (card_id) REFERENCES card (ID)
);

CREATE TABLE dictionary_to_library (
	ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
	dictionary_id INT UNSIGNED NOT NULL,
	library_id INT UNSIGNED NOT NULL,
	FOREIGN KEY (dictionary_id) REFERENCES dictionary (ID) ON DELETE CASCADE,
	FOREIGN KEY (library_id) REFERENCES cards_library (ID) ON DELETE CASCADE,
	UNIQUE(dictionary_id, library_id)
);


DROP PROCEDURE IF EXISTS borrow_dict;

DELIMITER $$
CREATE PROCEDURE borrow_dict(
	IN dict_id INT,
  IN thief_id INT
)
BEGIN
  DECLARE done BOOL DEFAULT FALSE;
  DECLARE cur_card_id INT;
  DECLARE cur_c_l_id INT;
  DECLARE new_dict_id, dict_owner INT;
  DECLARE dict_name VARCHAR(128);
  DECLARE dict_desc TEXT;
	DECLARE c1 CURSOR FOR
		SELECT card.ID
		FROM dictionary_to_library d_l
		JOIN cards_library c_l ON (d_l.library_id = c_l.id)
		JOIN card card ON(card.id = c_l.card_id)
		WHERE d_l.dictionary_id = dict_id;
	DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;
	SELECT d.name, d.description, d.UserID FROM dictionary d WHERE d.ID = dict_id INTO dict_name, dict_desc, dict_owner;
	INSERT INTO dictionary (name, description, UserID) VALUES (dict_name, dict_desc, thief_id);
	SELECT LAST_INSERT_ID() INTO new_dict_id;
	OPEN c1;

	read_loop: LOOP
		FETCH c1 INTO cur_card_id;
		IF done THEN
			LEAVE read_loop;
		END IF;
		INSERT INTO cards_library(card_id, count) VALUES (cur_card_id, 1);
		SELECT LAST_INSERT_ID() INTO cur_c_l_id;
		INSERT INTO dictionary_to_library(dictionary_id, library_id) VALUES (new_dict_id, cur_c_l_id);
	END LOOP;
	CLOSE c1;
	SELECT new_dict_id, dict_name, dict_desc, thief_id;
END $$
DELIMITER ;