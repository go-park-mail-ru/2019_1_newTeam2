package database

import (
	"crypto/sha256"
	"strconv"
	"github.com/user/2019_1_newTeam2/models"
)

type Database struct {
	// should be sqlx.db etc
	// now some map, where we can get users from
	Data   map[int]models.User
	LastId int
}

func NewDataBase() (*Database, error) {
	// error is possible error from database
	db := new(Database)
	// some code(real code working with db)

	data := make(map[int]models.User)
	LastId := 10
	h := sha256.New()
	h.Write([]byte("pass"))

	for i := 0; i < LastId; i++ {
		data[i] = models.User{i, "test_user_" + strconv.Itoa(i), "kek@lol.kl", string(h.Sum(nil)), 0, 1, 0, "files/avatars/1.jpg"}
	}
	db.Data = data
	return db, nil
}