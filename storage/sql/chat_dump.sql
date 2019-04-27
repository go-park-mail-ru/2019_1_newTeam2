USE chat_wordtrainer;

DROP TABLE IF EXISTS dialog;
DROP TABLE IF EXISTS broadcast_dialog;
DROP TABLE IF EXISTS message;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS word;

CREATE TABLE user (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    Username VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE message (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    data TEXT
);

CREATE TABLE dialog (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    MessageID INT UNSIGNED NOT NULL,
    User1 INT UNSIGNED NOT NULL,
    User2 INT UNSIGNED NOT NULL,
    FOREIGN KEY (MessageID) REFERENCES message (ID),
    FOREIGN KEY (User1) REFERENCES user (ID),
    FOREIGN KEY (User2) REFERENCES user (ID)
);

CREATE TABLE broadcast_dialog (
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    MessageID INT UNSIGNED NOT NULL,
    Author INT UNSIGNED NOT NULL,
    FOREIGN KEY (MessageID) REFERENCES message (ID),
    FOREIGN KEY (Author) REFERENCES user (ID)
);