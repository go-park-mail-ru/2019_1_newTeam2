package storage

import (
	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) GetMessagesBroadcast(page int, rowsNum int) ([]models.Message, error) {
	messages := make([]models.Message, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)
	rows, err := db.Conn.Query(GetMessages, rowsNum, offset)
	if err != nil {
		db.Logger.Log(err)
		return messages, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		i++
		message := models.Message{}
		err := rows.Scan(&message.ID, &message.Data)
		if err != nil {
			db.Logger.Log(err)
			return messages, err
		}
		messages = append(messages, message)
	}
	if i == 0 {
		db.Logger.Log("GetMessagesBroadcast: Empty query")
	}
	return messages, nil
}

func (db *Database) AddMessage(UserID int, message string) error {
	tx, err := db.Conn.Begin()
	result, err := tx.Exec(
		AddMessage,
		message,
		UserID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	result, err = tx.Exec(
		AddMessageToBroadcastDialog,
		lastID,
		UserID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}