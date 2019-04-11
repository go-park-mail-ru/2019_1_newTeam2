package storage

import "github.com/user/2019_1_newTeam2/models"

func (db *Database) GetCards(dictId int, page int, rowsNum int) ([]models.Card, bool, error) {
	cards := make([]models.Card, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)
	rows, err := db.Conn.Query(CardsPaginate, dictId, rowsNum, offset)
	// TODO(sergeychur): implement query
	if err != nil {
		return cards, false, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		i++
		card := models.Card{}
		err := rows.Scan(&card.ID, &card.Word.LanguageId, &card.Word.Name,
			&card.Translation.LanguageId, &card.Translation.Name)
		if err != nil{
			return cards, false, err
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
	err := row.Scan(&card.ID, &card.Word.Name, &card.Word.LanguageId, &card.Translation.Name, &card.Translation.LanguageId)
	if err != nil {
		return models.Card{}, false, err
	}
	return card, true, nil
}