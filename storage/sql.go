package storage

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	//  utils
	UseDB = "USE %s"

	//  user
	GetUserByUsernameQuery   = "SELECT ID, Username, Email, Password, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE Username = ?"
	GetUserByIDQuery         = "SELECT ID, Username, Email, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE ID = ?"
	AddUserQuery             = "INSERT INTO wordtrainer.user (Username, Email, Password, LangId, PronounceON, Score, AvatarPath) VALUES (?, ?, ?, ?, ?, ?, ?)"
	UpdateUserQuery          = "UPDATE wordtrainer.user SET Username = ?, Email = ?, LangId = ?, PronounceON = ? WHERE ID = ?"
	DeleteUserQuery          = "DELETE FROM wordtrainer.user WHERE ID = ?"
	UpdateImagePathUserQuery = "UPDATE wordtrainer.user SET AvatarPath = ? WHERE ID = ?"
	UsersPaginate            = "SELECT u.Username, u.Score " +
		"FROM wordtrainer.user u JOIN ( SELECT id FROM wordtrainer.user ORDER BY score " +
		"LIMIT ? OFFSET ?) l ON (u.id = l.id) " +
		"ORDER BY score;"

	// dictionary
	CreateEmptyDictionary = "INSERT INTO wordtrainer.dictionary (name, description, UserID) VALUES (?, ?, ?)"
	UpdateDictionary      = "UPDATE wordtrainer.dictionary SET name = ?, description = ? WHERE ID = ?"
	DeleteDictionary      = "DELETE FROM wordtrainer.dictionary WHERE ID = ?"
	GetDictById           = "SELECT id, name, description, UserId FROM wordtrainer.dictionary " +
		"WHERE id = ?"
	DictsPaginate = "SELECT d.ID, d.name, d.description, d.UserId " +
		"FROM wordtrainer.dictionary d JOIN ( SELECT id FROM wordtrainer.dictionary WHERE UserId = ? ORDER BY id " +
		"LIMIT ? OFFSET ?) l ON (d.id = l.id) " +
		"ORDER BY id;"
	GetDictOwner   = "SELECT UserId FROM wordtrainer.dictionary WHERE id = ?"
	DictBorrowProc = "CALL wordtrainer.borrow_dict(?, ?)"

	//  word
	GetWord = "SELECT ID FROM wordtrainer.word WHERE name = ? AND LangID = ?"

	//  card
	CreateCard    = "INSERT INTO wordtrainer.card (word, translation) VALUES (?, ?)"
	GetCard       = "SELECT ID FROM wordtrainer.card WHERE word = ? AND translation = ?"
	CardsPaginate = "SELECT c_l.id, w1.LangID, w1.name, w2.LangID, w2.name, c_l.guessed / c_l.seen FROM wordtrainer.card c " +
		"JOIN (SELECT wordtrainer.card.id " +
		"FROM wordtrainer.dictionary_to_library d_l " +
		"JOIN wordtrainer.cards_library c_l ON (d_l.library_id = c_l.id) " +
		"JOIN wordtrainer.card card ON(card.id = c_l.card_id) " +
		"WHERE d_l.dictionary_id = ? " +
		"ORDER BY card_id LIMIT ? OFFSET ?) l " +
		"ON (c.id = l.id) " +
		"JOIN wordtrainer.cards_library c_l ON ( c_l.card_id = c.id) " +
		"JOIN wordtrainer.word w1 on (w1.id = c.word) " +
		"JOIN wordtrainer.word w2 on (w2.id = c.translation) " +
		"ORDER BY id DESC"
	GetCardById = "SELECT c_l.id, w1.name, w1.LangID, " +
		" w2.name, w2.LangID, c_l.guessed / c_l.seen from wordtrainer.cards_library c_l " +
		"join wordtrainer.card c on c_l.card_id = c.id " +
		" join wordtrainer.word w1 " +
		"on (w1.id = c.word) join wordtrainer.word w2 on " +
		"(w2.id = c.translation) where c_l.id = ?"
	CardsForGame = "SELECT c_l.id, w1.name, w2.name FROM wordtrainer.dictionary_to_library d_l " +
		"JOIN wordtrainer.cards_library c_l ON (c_l.id = d_l.library_id )" +
		"JOIN wordtrainer.card card ON(card.id = c_l.card_id) " +
		"JOIN wordtrainer.word w1 on (w1.id = card.word) " +
		"JOIN wordtrainer.word w2 on (w2.id = card.translation) " +
		"WHERE d_l.dictionary_id = ? " +
		"ORDER BY c_l.guessed / c_l.seen ASC LIMIT ?"
	/*CountCardsInDict = "SELECT count(*) FROM dictionary_to_library d_l " +
		"JOIN cards_library c_l ON (c_l.id = d_l.library_id) " +
		"WHERE d_l.dictionary_id = ?"*/
	GetWordsFromDict = "select w.name from dictionary_to_library d_l " +
		"join cards_library c_l on(c_l.id=d_l.library_id) join card c on(c.id =  c_l.card_id) " +
		"join word w on (c.translation = w.id) WHERE d_l.dictionary_id = ? ORDER BY c_l.guessed / c_l.seen ASC LIMIT ?"

	TriggerDeleteCard = "DELETE FROM wordtrainer.cards_library WHERE ID IN ( SELECT library_id FROM wordtrainer.dictionary_to_library WHERE dictionary_id = ?)"

	//  language
	CreateLanguage = "INSERT INTO wordtrainer.language (name) VALUES (?)"
	GetLangs       = "SELECT * FROM wordtrainer.language"
	GetLangById    = "SELECT * FROM wordtrainer.language WHERE name = ?"

	//  cards_library
	CreateCardsLibrary     = "INSERT INTO wordtrainer.cards_library (card_id, count) VALUES (?, ?)"
	DeleteListCardsLibrary = "DELETE FROM wordtrainer.cards_library WHERE ID in ?"
	GetIDCardsLibrary      = "SELECT ID FROM wordtrainer.cards_library WHERE card_id = ?"
	IncrCountCardsLibrary  = "UPDATE wordtrainer.cards_library SET count = count + 1 WHERE ID = ?"
	DecrCountCardsLibrary  = "UPDATE wordtrainer.cards_library SET count = count - 1 WHERE ID = ?"

	// mb change, talk about it
	UpdateFrequency = "UPDATE wordtrainer.cards_library SET seen = if(if_seen, seen + 1, seen), if_seen = true, guessed = guessed + ? where id = ?"


	//  dictionary_to_library
	CreateDictionaryToLibrary     = "INSERT INTO wordtrainer.dictionary_to_library (dictionary_id, library_id) VALUES (?, ?)"
	DeleteDictionaryToLibraryByID = "DELETE FROM wordtrainer.dictionary_to_library WHERE dictionary_id = ? AND library_id = ?"
	GetLibraryIDByDictionaryID    = "SELECT library_id FROM wordtrainer.dictionary_to_library WHERE dictionary_id = ?"

	//  chat
	GetMessages = "SELECT m.UserId, m.data FROM chat_wordtrainer.broadcast_dialog JOIN chat_wordtrainer.message m ON (MessageId = m.ID) LIMIT ? OFFSET ?";
	AddMessage = "INSERT INTO chat_wordtrainer.message (data, UserId) VALUES (?, ?)";
	AddMessageToBroadcastDialog = "INSERT INTO chat_wordtrainer.broadcast_dialog (MessageId, UserId) VALUES(?, ?)"
)
