package storage

import (
	"github.com/user/2019_1_newTeam2/models"
	"math/rand"
)

func Find(arr []int, num int) bool {
	for _, i := range arr {
		if i == num {
			return true
		}
	}
	return false
}

func RandNum(max int, count int) []int {
	arr := make([]int, count)
	for i := range arr {
		res := rand.Int() % max + 1
		for Find(arr, res) == true {
			res = rand.Int() % max + 1
		}
		arr[i] = res
	}
	return arr
}

func (db *Database) CreateTask() (models.GameQuestion, error) {
	var LastID int
	res := db.Conn.QueryRow(GameGetLastId)
	err := res.Scan(&LastID)
	if err != nil {
		db.Logger.Log("CreateTask", err)
		return models.GameQuestion{}, err
	}
	task := models.GameQuestion{}
	RandInt := RandNum(LastID,4)
	AnswerId := rand.Int() % 4

	res = db.Conn.QueryRow(GameGetTranslate, RandInt[AnswerId])
	err = res.Scan(&task.Question)
	if err != nil {
		db.Logger.Log("CreateTask", err)
		return models.GameQuestion{}, err
	}
	for i := range RandInt {
		res = db.Conn.QueryRow(GameGetWord, RandInt[i])
		err = res.Scan(&task.Words[i])
		if err != nil {
			db.Logger.Log("CreateTask", err)
			return models.GameQuestion{}, err
		}
	}
	task.Answer = task.Words[AnswerId]
	return task, nil
}