package storage

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	UseDB = "USE wordtrainer"
	GetUserByUsernameQuery = "SELECT ID, Username, Email, Password, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE Username = ?"
	GetUserByIDQuery = "SELECT ID, Username, Email, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE ID = ?"
	AddUserQuery = "INSERT INTO user (Username, Email, Password, LangId, PronounceON, Score, AvatarPath) VALUES (?, ?, ?, ?, ?, ?, ?)"
	UpdateUserQuery = "UPDATE user SET Username = ?, Email = ?, Password = ?, LangId = ?, PronounceON = ?, WHERE ID = ?"
	DeleteUserQuery = "DELETE FROM user WHERE ID = ?"
	UpdateImagePathUserQuery = "UPDATE user SET AvatarPath = ? WHERE ID = ?"
	GetLangs = "SELECT * FROM language"
	UsersPaginate = "SELECT u.Username, u.Score " +
		"FROM user u JOIN ( SELECT id FROM user ORDER BY score " +
		"LIMIT ? OFFSET ?) l ON (u.id = l.id) " +
		"ORDER BY score;"
	CardsPaginate = ""
	GetCardById = "SELECT c.id, w1.name, w1.LangID, " +
		" w2.name, w2.LangID from card c join word w1 " +
		"on (w1.id = c.word) join word w2 on " +
		"(w2.id = c.translation) where c.id = ?"
	GetDictById = "SELECT id, name, description, userId FROM dictionary " +
				"WHERE id = ?"
	DictsPaginate = "SELECT d.ID, d.name, d.description, d.UserId " +
		"FROM dictionary d JOIN ( SELECT id FROM dictionary WHERE UserId = ? ORDER BY id " +
		"LIMIT ? OFFSET ?) l ON (d.id = l.id) " +
		"ORDER BY id;"
)

