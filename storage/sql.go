package storage

import (
	_ "github.com/go-sql-driver/mysql"
)

var GetUserByUsernameQuery = "SELECT ID, Username, Email, Password, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE Username = ?"
var GetUserByIDQuery = "SELECT ID, Username, Email, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE ID = ?"
var AddUserQuery = "INSERT INTO user (Username, Email, Password, LangId, PronounceON, Score, AvatarPath) VALUES (?, ?, ?, ?, ?, ?, ?)"
var UpdateUserQuery = "UPDATE user SET Username = ?, Email = ?, Password = ?, LangId = ?, PronounceON = ?, WHERE ID = ?"
var DeleteUserQuery = "DELETE FROM user WHERE ID = ?"
var UpdateImagePathUserQuery = "UPDATE user SET AvatarPath = ? WHERE ID = ?"
