package storage

import (
	"fmt"
)

func (db *Database) CreateCardsLibrary(CardID int) (int, error) {
	result, CreateErr := db.Conn.Exec(
		CreateCardsLibrary,
		0,
		CardID,
		1,
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
