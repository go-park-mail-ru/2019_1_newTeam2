package storage

import (
	_ "github.com/go-sql-driver/mysql"
)

//  user
var GetUserByUsernameQuery = "SELECT ID, Username, Email, Password, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE Username = ?"
var GetUserByIDQuery = "SELECT ID, Username, Email, LangID, PronounceON, Score, AvatarPath FROM wordtrainer.user WHERE ID = ?"
var AddUserQuery = "INSERT INTO user (Username, Email, Password, LangId, PronounceON, Score, AvatarPath) VALUES (?, ?, ?, ?, ?, ?, ?)"
var UpdateUserQuery = "UPDATE user SET Username = ?, Email = ?, Password = ?, LangId = ?, PronounceON = ?, WHERE ID = ?"
var DeleteUserQuery = "DELETE FROM user WHERE ID = ?"
var UpdateImagePathUserQuery = "UPDATE user SET AvatarPath = ? WHERE ID = ?"

// dictionary
var CreateEmptyDictionary = "INSERT INTO dictionary (name, description, UserID) VALUES (?, ?, ?)"
var UpdateDictionary = "UPDATE dictionary SET name = ?, description = ? WHERE ID = ?"
var DeleteDictionary = "DELETE FROM dictionary_to_library WHERE dictionary_id = ?"

//  word
var CreateWord = "INSERT INTO word (name, LangID) VALUES (?, ?)"

//  card
var CreateCard = "INSERT INTO card (word, translation) VALUES (?, ?)"

//  language
var CreateLanguage = "INSERT INTO language (name) VALUES (?)"

//  cards_library
var CreateCardsLibrary = "INSERT INTO cards_library (frequency, card_id) VALUES (?, ?)"
var DeleteListCardsLibrary = "DELETE FROM cards_library WHERE ID in ?"

//  dictionary_to_library
var CreateDictionaryToLibrary = "INSERT INTO dictionary_to_library (dictionary_id, library_id) VALUES (?, ?)"
