package storage

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	GetUserByUsernameQuery = "SELECT ID, Username, Email, Password, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE Username = ?"
	GetUserByIDQuery = "SELECT ID, Username, Email, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE ID = ?"
	AddUserQuery = "INSERT INTO user (Username, Email, Password, LangId, PronounceON, Score, AvatarPath) VALUES (?, ?, ?, ?, ?, ?, ?)"
	UpdateUserQuery = "UPDATE user SET Username = ?, Email = ?, Password = ?, LangId = ?, PronounceON = ?, WHERE ID = ?"
	DeleteUserQuery = "DELETE FROM user WHERE ID = ?"
	UpdateImagePathUserQuery = "UPDATE user SET AvatarPath = ? WHERE ID = ?"
	GetLangs = "SELECT * FROM language"
)