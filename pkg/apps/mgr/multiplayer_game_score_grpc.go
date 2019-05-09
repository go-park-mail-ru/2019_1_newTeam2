package mgr

import (
	"context"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
	"strconv"
)

type UserScoreUpdaterManager struct {
	DB interfaces.DBMGRInterface
}

func NewUserScoreUpdaterManager(DBUser string, DBPassUser string) *UserScoreUpdaterManager {
	server := UserScoreUpdaterManager{}
	newDB, err := storage.NewDataBase(DBUser, DBPassUser)
	if err != nil {
		return nil
	}
	server.DB = newDB
	return &server
}

func (usum *UserScoreUpdaterManager) UpdateUserScore(ctx context.Context, in *UserScore) (*NewScore, error) {
	user, err := usum.DB.UpdateScore(in.Username, int(in.AddScore))
	if err != nil {
		return &NewScore{}, err
	}
	score := strconv.Itoa(user.Score)
	return &NewScore{Score: score}, nil
}
