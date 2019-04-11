package storage

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	//  utils
	UseDB = "USE wordtrainer"

	//  user
	GetUserByUsernameQuery   = "SELECT ID, Username, Email, Password, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE Username = ?"
	GetUserByIDQuery         = "SELECT ID, Username, Email, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE ID = ?"
	AddUserQuery             = "INSERT INTO user (Username, Email, Password, LangId, PronounceON, Score, AvatarPath) VALUES (?, ?, ?, ?, ?, ?, ?)"
	UpdateUserQuery          = "UPDATE user SET Username = ?, Email = ?, Password = ?, LangId = ?, PronounceON = ?, WHERE ID = ?"
	DeleteUserQuery          = "DELETE FROM user WHERE ID = ?"
	UpdateImagePathUserQuery = "UPDATE user SET AvatarPath = ? WHERE ID = ?"
	UsersPaginate            = "SELECT u.Username, u.Score " +
		"FROM user u JOIN ( SELECT id FROM user ORDER BY score " +
		"LIMIT ? OFFSET ?) l ON (u.id = l.id) " +
		"ORDER BY score;"

	// dictionary
	CreateEmptyDictionary = "INSERT INTO dictionary (name, description, UserID) VALUES (?, ?, ?)"
	UpdateDictionary      = "UPDATE dictionary SET name = ?, description = ? WHERE ID = ?"
	DeleteDictionary      = "DELETE FROM dictionary_to_library WHERE dictionary_id = ?"
	GetDictById           = "SELECT id, name, description, userId FROM dictionary " +
		"WHERE id = ?"
	DictsPaginate = "SELECT d.ID, d.name, d.description, d.UserId " +
		"FROM dictionary d JOIN ( SELECT id FROM dictionary WHERE UserId = ? ORDER BY id " +
		"LIMIT ? OFFSET ?) l ON (d.id = l.id) " +
		"ORDER BY id;"

	//  word
	GetWord = "SELECT ID FROM word WHERE name = ? AND LangID = ?"

	//  card
	CreateCard = "INSERT INTO card (word, translation) VALUES (?, ?)"

	CardsPaginate = "SELECT c.id, w1.LangID, w1.name, w2.LangID, w2.name FROM card c JOIN " +
	"(SELECT card.id " +
	"FROM dictionary_to_library d_l " +
	"JOIN cards_library c_l ON (d_l.library_id = c_l.id) " +
	"JOIN card card ON(card.id = c_l.card_id) " +
	"WHERE d_l.dictionary_id = ? " +
	"ORDER BY card_id LIMIT ? OFFSET ?) l " +
	"ON (c.id = l.id) " +
	"JOIN word w1 on (w1.id = c.word) " +
	"JOIN word w2 on (w2.id = c.translation) " +
	"ORDER BY id"
	GetCardById   = "SELECT c.id, w1.name, w1.LangID, " +
		" w2.name, w2.LangID from card c join word w1 " +
		"on (w1.id = c.word) join word w2 on " +
		"(w2.id = c.translation) where c.id = ?"

	//  language
	CreateLanguage = "INSERT INTO language (name) VALUES (?)"
	GetLangs       = "SELECT * FROM language"

	//  cards_library
	CreateCardsLibrary     = "INSERT INTO cards_library (frequency, card_id) VALUES (?, ?)"
	DeleteListCardsLibrary = "DELETE FROM cards_library WHERE ID in ?"

	//  dictionary_to_library
	CreateDictionaryToLibrary = "INSERT INTO dictionary_to_library (dictionary_id, library_id) VALUES (?, ?)"
)
