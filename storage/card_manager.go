package storage

import "fmt"

func (db *Database) CreateCard(WordID int, TranslationID int) (int, error) {
	result, CreateErr := db.Conn.Exec(
		CreateCard,
		WordID,
		TranslationID,
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
