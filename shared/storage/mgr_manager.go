package storage

import (
	"github.com/user/2019_1_newTeam2/shared/models"
)

func (db *Database) UpdateScore(username string, added_points int) (models.UserScore, error) {
	user := models.UserScore{}
	user.Username = username
	row := db.Conn.QueryRow(GetUserScore, username)
	err := row.Scan(&user.Score)
	user.Score = user.Score + added_points
	if err != nil {
		db.Logger.Log("UpdateScore: ", err)
		return models.UserScore{}, err
	}
	_, err = db.Conn.Exec(
		UpdateUserScore,
		user.Score,
		user.Username,
	)
	if err != nil {
		db.Logger.Log("UpdateScore: ", err)
		return models.UserScore{}, err
	}
	return user, nil
}
