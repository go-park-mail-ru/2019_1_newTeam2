package storage

import "fmt"

func (db *Database) AddToDictionaryToLibrary(lastID int, CardsLibraryID int) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return fmt.Errorf("AddToDictionaryToLibrary: transaction err")
	}
	_, CreateErr := tx.Exec(
		CreateDictionaryToLibrary,
		lastID,
		CardsLibraryID,
	)

	if CreateErr != nil {
		_ = tx.Rollback()
		return fmt.Errorf("AddToDictionaryToLibrary: word not create")
	}
	_ = tx.Commit()
	return nil
}

func (db *Database) DeleteDictionaryToLibraryByID(dictionaryID int, cardID int) error {
	_, DeleteErr := db.Conn.Exec(
		DeleteDictionaryToLibraryByID,
		dictionaryID,
		cardID,
	)

	if DeleteErr != nil {
		return DeleteErr
	}

	return nil
}
