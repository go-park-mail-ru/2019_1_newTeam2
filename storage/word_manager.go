package storage

import (
	"fmt"
	"strconv"

	"github.com/user/2019_1_newTeam2/models"
)

func CreateWord(db *Database, word *models.Word) (int, error) {
	Query := "INSERT INTO word (name, LangID) SELECT * FROM (SELECT \"" +
		word.Name + "\", " + strconv.Itoa(word.LanguageId) +
		") AS tmp WHERE NOT EXISTS " +
		"(SELECT name, LangID FROM word WHERE name = ? AND LangID = ?) LIMIT 1"
	result, CreateErr := db.Conn.Exec(
		Query,
		word.Name,
		word.LanguageId,
	)
	if CreateErr != nil {
		return 0, fmt.Errorf("Err: word not create and not found")
	}

	lastID, GetIDErr := result.LastInsertId()
	if GetIDErr != nil {
		return 0, fmt.Errorf("GetIDErr: can`t get last word id")
	}

	if lastID != 0 {
		return int(lastID), nil
	}

	var ID int64
	row := db.Conn.QueryRow(GetWord, word.Name, word.LanguageId)
	err := row.Scan(&ID)
	if err != nil {
		return 0, fmt.Errorf("Err: word not create and not found")
	}
	return int(ID), nil
}
