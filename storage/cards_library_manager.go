package storage

import (
	"fmt"

	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) DeleteCardInDictionary(dictionaryID int, cardID int) error {
	if err := db.DecrementCount(cardID); err != nil {
		db.Logger.Log(err)
		return err
	}
	if err := db.DeleteDictionaryToLibraryByID(dictionaryID, cardID); err != nil {
		db.Logger.Log(err)
		return err
	}
	return nil
}

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

func (db *Database) IfExistCardLibrary(CardID int) int {
	var ID int
	row := db.Conn.QueryRow(GetIDCardsLibrary, CardID)
	err := row.Scan(&ID)
	if err != nil {
		return 0
	}
	return ID
}

func (db *Database) IncrementCount(CardID int) error {
	_, err := db.Conn.Exec(
		IncrCountCardsLibrary,
		CardID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (db *Database) DecrementCount(CardID int) error {
	_, err := db.Conn.Exec(
		DectCountCardsLibrary,
		CardID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (db *Database) SetCardToDictionary(dictID int, card models.Card) error {
	var WordID, TranslationID, CardID, CardsLibraryID int
	var err error
	WordID, err = CreateWord(db, card.Word)
	if err != nil {
		db.Logger.Log(err)
		return err
	}
	TranslationID, err = CreateWord(db, card.Translation)
	if err != nil {
		db.Logger.Log(err)
		return err
	}
	CardID, err = db.CreateCard(WordID, TranslationID)
	if err != nil {
		db.Logger.Log(err)
		return err
	}
	if CardsLibraryID = db.IfExistCardLibrary(CardID); CardsLibraryID == 0 {
		CardsLibraryID, err = db.CreateCardsLibrary(CardID)
		if err != nil {
			db.Logger.Log(err)
			return err
		}
	} else {
		if err = db.IncrementCount(CardsLibraryID); err != nil {
			db.Logger.Log(err)
			return err
		}
	}
	err = db.AddToDictionaryToLibrary(dictID, CardsLibraryID)
	if err != nil {
		db.Logger.Log(err)
		return err
	}
	return nil
}
