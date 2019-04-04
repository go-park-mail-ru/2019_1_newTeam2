package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var createTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS wordtrainer DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE wordtrainer;`,
	`CREATE TABLE IF NOT EXISTS user (
		ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		Username VARCHAR(255) NOT NULL,
		Email VARCHAR(255) NOT NULL,
		Password VARCHAR(255) NOT NULL,
		LangID INT UNSIGNED NOT NULL,
		PronounceON TINYINT NOT NULL,
		Score INT UNSIGNED NOT NULL,
		AvatarPath VARCHAR(255) NOT NULL
	);`,
}

var GetUserByIDQuery = "SELECT ID, Username, Email, Password, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE ID = ?"

func createTable(conn *sql.DB) error {
	for _, stmt := range createTableStatements {
		_, err := conn.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
