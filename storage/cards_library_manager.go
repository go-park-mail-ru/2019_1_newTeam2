package storage

import (
	"fmt"

	"github.com/user/2019_1_newTeam2/models"
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

func (db *Database) SetCardToDictionary(dictID int, card models.Card) error {
	var WordID, TranslationID, CardID, CardsLibraryID int
	var err error
	WordID, err = CreateWord(db, card.Word)
	if err != nil {
		fmt.Println(err)
		return err
	}
	TranslationID, err = CreateWord(db, card.Translation)
	if err != nil {
		fmt.Println(err)
		return err
	}
	CardID, err = db.CreateCard(WordID, TranslationID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	CardsLibraryID, err = db.CreateCardsLibrary(CardID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = db.AddToDictionaryToLibrary(dictID, CardsLibraryID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
