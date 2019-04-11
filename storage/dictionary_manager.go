package storage

import (
	"fmt"

	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) DictionaryDelete(DictID int) error {
	cards_id := []int{}
	rows, _ := db.Conn.Query("SELECT library_id FROM dictionary_to_library WHERE dictionary_id = ?", DictID)
	for rows.Next() {
		var card_id int
		err := rows.Scan(card_id)
		if err != nil {
			return err
		}
		cards_id = append(cards_id, card_id)
	}

	return nil
}

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

func (db *Database) DictionaryCreate(UserID int, Name string, Description string, Cards []models.Card) error {
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
		var err error
		WordID, err = CreateWord(db, it.Word)
		if err != nil {
			return err
		}
		TranslationID, err = CreateWord(db, it.Translation)
		if err != nil {
			return err
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

func (db *Database) GetDicts(userId int, page int, rowsNum int) ([]models.DictionaryInfo, bool, error) {
	dicts := make([]models.DictionaryInfo, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)
	rows, err := db.Conn.Query(DictsPaginate, userId, rowsNum, offset)
	if err != nil {
		return dicts, false, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		i++
		dict := models.DictionaryInfo{}
		err := rows.Scan(&dict.ID, &dict.Name, &dict.Description /*, &dict.UserId*/)
		// TODO(sergeychur): say about userId, may be useful, if no delete
		if err != nil {
			return dicts, false, err
		}
		dicts = append(dicts, dict)
	}
	if i == 0 {
		return dicts, false, nil
	}
	return dicts, true, nil
}
func (db *Database) GetDict(dictId int) (models.DictionaryInfo, bool, error) {
	dict := models.DictionaryInfo{}
	row := db.Conn.QueryRow(GetDictById, dictId)
	err := row.Scan(&dict.ID, &dict.Name, &dict.Description /*, &dict.UserId*/)
	if err != nil {
		return models.DictionaryInfo{}, false, err
	}
	return dict, true, nil
}
