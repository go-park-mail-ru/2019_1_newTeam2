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
		var err error
		WordID, err = db.CreateWord(it.Word)
		if err != nil {
			return err
		}
		TranslationID, err = db.CreateWord(it.Translation)
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

// УДАЛИТЬ

// func (db *Database) GetCard(cardId int) (models.AddedToDictCard, bool, error) {
// 	card := models.AddedToDictCard{}
// 	card.Word = new(models.AddedToDictWord)
// 	card.Translation = new(models.AddedToDictWord)

// 	const GetCardById = "SELECT c.id, w1.name, w1.LangID, " +
// 		" w2.name, w2.LangID from card c join word w1 " +
// 		"on (w1.id = c.word) join word w2 on " +
// 		"(w2.id = c.translation) where c.id = ?"

// 	row := db.Conn.QueryRow(GetCardById, cardId)
// 	err := row.Scan(&card.ID, &card.Word.Name, &card.Word.Language, &card.Translation.Name, &card.Translation.Language)
// 	if err != nil {
// 		return models.AddedToDictCard{}, false, err
// 	}
// 	return card, true, nil
// }
