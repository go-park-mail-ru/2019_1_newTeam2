package storage

import (
	"github.com/user/2019_1_newTeam2/shared/models"
)

func (db *Database) GetMessagesBroadcast(page int, rowsNum int) (models.Messages, error) {
	messages := make([]models.Message, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)
	rows, err := db.Conn.Query(GetMessages, rowsNum, offset)
	if err != nil {
		db.Logger.Log(err)
		return messages, err
	}
	defer func() {
		_ = rows.Close()
	}()
	i := 0
	for rows.Next() {
		i++
		message := models.Message{}
		err := rows.Scan(&message.ID, &message.UserName, &message.Data)
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

func (db *Database) AddMessage(UserName string, UserID int, message string) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec(
		AddMessage,
		message,
		UserID,
		UserName,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	_, err = tx.Exec(
		AddMessageToBroadcastDialog,
		lastID,
		UserID,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	_ = tx.Commit()
	return nil
}
