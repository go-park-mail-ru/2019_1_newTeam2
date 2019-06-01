package storage

import (
	"database/sql"
	"fmt"

	"github.com/user/2019_1_newTeam2/shared/models"
)

func (db *Database) DeleteCardInDictionary(userId, dictionaryID int, cardID int) error {
	ifOwner := false
	err := db.Conn.QueryRow(CheckOwner, dictionaryID, userId).Scan(&ifOwner)
	if err != nil {
		return err
	}
	if !ifOwner {
		return fmt.Errorf("u cheater")
	}
	if err = db.DecrementCount(cardID); err != nil {
		db.Logger.Log(err)
		return err
	}
	if err = db.DeleteDictionaryToLibraryByID(dictionaryID, cardID); err != nil {
		db.Logger.Log(err)
		return err
	}
	return nil
}

func (db *Database) CreateCardsLibrary(CardID int) (int, error) {
	tx, err := db.Conn.Begin()
	if err != nil {
		db.Logger.Log("CreateCardsLibrary: transaction error")
		return 0, err
	}

	result, CreateErr := tx.Exec(
		CreateCardsLibrary,
		CardID,
		1,
	)
	if CreateErr != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("CreateErr: word not create")
	}
	lastID, GetIDErr := result.LastInsertId()
	if GetIDErr != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("GetIDErr: can`t get last dict id")
	}
	_ = tx.Commit()
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
		DecrCountCardsLibrary,
		CardID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (db *Database) SetCardToDictionary(userId int, dictID int, card models.Card) error {
	ifOwner := false
	err := db.Conn.QueryRow(CheckOwner, dictID, userId).Scan(&ifOwner)
	if err == sql.ErrNoRows {
		return fmt.Errorf("u cheater")
	}
	if err != nil {
		return err
	}
	if !ifOwner {
		return fmt.Errorf("u cheater")
	}
	var WordID, TranslationID, CardID, CardsLibraryID int
	//var err error
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
