USE chat_wordtrainer;

DROP TABLE IF EXISTS russian_english_words;
DROP TABLE IF EXISTS message_in_dialog;
DROP TABLE IF EXISTS dialog;
DROP TABLE IF EXISTS broadcast_dialog;
DROP TABLE IF EXISTS message;

CREATE TABLE russian_english_words (
    Word VARCHAR(256) PRIMARY KEY,
    Translate VARCHAR(256)
);

CREATE TABLE message (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    data TEXT,
    UserId INT UNSIGNED NOT NULL
);

CREATE TABLE dialog (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    MessageId INT UNSIGNED NOT NULL,
    FOREIGN KEY (MessageID) REFERENCES message (ID)
);

CREATE TABLE message_in_dialog (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    DialogId INT UNSIGNED NOT NULL,
    MessageId INT UNSIGNED NOT NULL,
    FOREIGN KEY (DialogId) REFERENCES dialog (ID),
    FOREIGN KEY (MessageId) REFERENCES message (ID)
);

CREATE TABLE broadcast_dialog (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    MessageId INT UNSIGNED NOT NULL,
    UserId INT UNSIGNED NOT NULL,
    FOREIGN KEY (MessageId) REFERENCES message (ID)
);