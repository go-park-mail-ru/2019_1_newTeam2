package storage

import (
	"fmt"

	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) WordIfExist(word *models.AddedToDictWord) (int, bool) {

	return 0, false
}

func (db *Database) CreateWord(word *models.AddedToDictWord) (int, error) {
	result, CreateErr := db.Conn.Exec(
		CreateWord,
		word.Name,
		word.Language,
	)

	if CreateErr != nil {
		return 0, fmt.Errorf("CreateErr: word not create")
	}

	lastID, GetIDErr := result.LastInsertId()
	if GetIDErr != nil {
		return 0, fmt.Errorf("GetIDErr: can`t get last dict id")
	}
	return int(lastID), nil
}
