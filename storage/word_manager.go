package storage

import (
	"fmt"
	"strconv"

	"github.com/user/2019_1_newTeam2/models"
)

<<<<<<< HEAD
func (db *Database) WordIfExist(word *models.Word) (int, bool) {

	return 0, false
}

func (db *Database) CreateWord(word *models.Word) (int, error) {
=======
func (db *Database) CreateWord(word *models.AddedToDictWord) (int, error) {
	Query := "INSERT INTO word (name, LangID) SELECT * FROM (SELECT \"" + word.Name + "\", " + strconv.Itoa(word.Language) + ") AS tmp WHERE NOT EXISTS (SELECT name, LangID FROM word WHERE name = ? AND LangID = ?) LIMIT 1"
>>>>>>> 02588008e9e7159ee32237aaf80dae5024859e44
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
		return 0, fmt.Errorf("GetIDErr: can`t get last dict id")
	}

	if lastID != 0 {
		return int(lastID), nil
	}

	var ID int64
	row := db.Conn.QueryRow(GetWord, word.Name, word.Language)
	err := row.Scan(&ID)
	if err != nil {
		return 0, fmt.Errorf("Err: word not create and not found")
	}
	return int(ID), nil
}
