package storage

import (
	"fmt"
	"strconv"

	"github.com/user/2019_1_newTeam2/shared/models"
)

func CreateWord(db *Database, word *models.Word) (int, error) {
	Query := "INSERT INTO wordtrainer.word (name, LangID) SELECT * FROM (SELECT \"" +
		word.Name + "\", " + strconv.Itoa(word.LanguageId) +
		") AS tmp WHERE NOT EXISTS " +
		"(SELECT name, LangID FROM wordtrainer.word WHERE name = ? AND LangID = ?) LIMIT 1"
	tx, err := db.Conn.Begin()
	if err != nil {
		return 0, fmt.Errorf("transaction errpr")
	}
	result, CreateErr := tx.Exec(
		Query,
		word.Name,
		word.LanguageId,
	)
	if CreateErr != nil {
		_ = tx.Rollback()
		fmt.Println("Error: ", err)
		return 0, fmt.Errorf("Err: word not create and not found")
	}

	lastID, GetIDErr := result.LastInsertId()
	if GetIDErr != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("GetIDErr: can`t get last word id")
	}

	_ = tx.Commit()
	if lastID != 0 {
		return int(lastID), nil
	}

	tx, err = db.Conn.Begin()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	var ID int64
	row := tx.QueryRow(GetWord, word.Name, word.LanguageId)
	err = row.Scan(&ID)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println("Error: ", err)
		return 0, fmt.Errorf("Err: word not create and not found")
	}
	_ = tx.Commit()
	return int(ID), nil
}
