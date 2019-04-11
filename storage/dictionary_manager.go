package storage

import (
	"fmt"

	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) DictionaryUpdate(DictID int, Name string, Description string) error {
	_, UpdateErr := db.Conn.Exec(
		UpdateDictionary,
		Name,
		Description,
		DictID,
	)
	if UpdateErr != nil {
		fmt.Println("UpdateErr: user not update: ", UpdateErr)
		return fmt.Errorf("UpdateErr: user not update")
	}

	return nil
}

func (db *Database) DictionaryCreate(UserID int, Name string, Description string, Cards []models.AddedToDictCard) error {
	result, CreateErr := db.Conn.Exec(
		CreateEmptyDictionary,
		Name,
		Description,
		UserID,
	)
	if CreateErr != nil {
		fmt.Println("CreateErr: user not create: ", CreateErr)
		return fmt.Errorf("CreateErr: user not create")
	}

	lastID, GetIDErr := result.LastInsertId()
	if GetIDErr != nil {
		fmt.Println("GetIDErr: can`t get last dict id")
		return fmt.Errorf("GetIDErr: can`t get last dict id")
	}
	for _, it := range Cards {
		var WordID, TranslationID, CardID, CardsLibraryID int
		var res bool
		var err error
		WordID, res = db.WordIfExist(it.Word)
		if !res {
			WordID, err = db.CreateWord(it.Word)
			if err != nil {
				return err
			}
		}
		TranslationID, res = db.WordIfExist(it.Translation)
		if !res {
			TranslationID, err = db.CreateWord(it.Translation)
			if err != nil {
				return err
			}
		}
		CardID, err = db.CreateCard(WordID, TranslationID)
		if err != nil {
			return err
		}
		CardsLibraryID, err = db.CreateCardsLibrary(CardID)
		if err != nil {
			return err
		}
		err = db.AddToDictionaryToLibrary(int(lastID), CardsLibraryID)
		if err != nil {
			return err
		}
	}
	return nil
}
