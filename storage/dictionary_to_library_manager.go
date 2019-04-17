package storage

import "fmt"

func (db *Database) AddToDictionaryToLibrary(lastID int, CardsLibraryID int) error {
	_, CreateErr := db.Conn.Exec(
		CreateDictionaryToLibrary,
		lastID,
		CardsLibraryID,
	)

	if CreateErr != nil {
		return fmt.Errorf("CreateErr: word not create")
	}

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
