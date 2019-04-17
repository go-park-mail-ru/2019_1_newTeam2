package storage

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) CreateCard(WordID int, TranslationID int) (int, error) {
	Query := "INSERT INTO wordtrainer.card (word, translation) SELECT * FROM (SELECT \"" +
		strconv.Itoa(WordID) + "\", " + strconv.Itoa(TranslationID) +
		") AS tmp WHERE NOT EXISTS " +
		"(SELECT word, translation FROM wordtrainer.card WHERE word = ? AND translation = ?) LIMIT 1"
	result, CreateErr := db.Conn.Exec(
		Query,
		WordID,
		TranslationID,
	)

	if CreateErr != nil {
		return 0, fmt.Errorf("Err: card not create and not found")
	}

	lastID, GetIDErr := result.LastInsertId()
	if GetIDErr != nil {
		return 0, fmt.Errorf("GetIDErr: can`t get last card id")
	}

	if lastID != 0 {
		return int(lastID), nil
	}

	var ID int64
	row := db.Conn.QueryRow(GetCard, WordID, TranslationID)
	err := row.Scan(&ID)
	if err != nil {
		return 0, fmt.Errorf("Err: card not create and not found")
	}
	return int(ID), nil
}

func (db *Database) GetCards(dictId int, page int, rowsNum int) ([]models.Card, bool, error) {
	cards := make([]models.Card, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)
	rows, err := db.Conn.Query(CardsPaginate, dictId, rowsNum, offset)
	if err != nil {
		db.Logger.Log(err)
		return cards, false, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		i++
		card := models.Card{}
		card.Word = new(models.Word)
		card.Translation = new(models.Word)
		err := rows.Scan(&card.ID, &card.Word.LanguageId, &card.Word.Name,
			&card.Translation.LanguageId, &card.Translation.Name, &card.Frequency)
		if err != nil {
			return cards, false, err
			db.Logger.Log(err)
		}
		cards = append(cards, card)
	}
	if i == 0 {
		return cards, false, nil
	}
	return cards, true, nil
}

func (db *Database) GetCard(cardId int) (models.Card, bool, error) {
	card := models.Card{}
	card.Word = new(models.Word)
	card.Translation = new(models.Word)
	row := db.Conn.QueryRow(GetCardById, cardId)
	err := row.Scan(&card.ID, &card.Word.Name, &card.Word.LanguageId,
		&card.Translation.Name, &card.Translation.LanguageId, &card.Frequency)
	if err == sql.ErrNoRows {
		return models.Card{}, false, nil
	}
	if err != nil {
		return models.Card{}, false, err
	}
	return card, true, nil
}
