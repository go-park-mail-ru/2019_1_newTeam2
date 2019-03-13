package database

import (
	"crypto/sha256"
	"strconv"

	"github.com/user/2019_1_newTeam2/models"
)

type Database struct {
	// should be sqlx.db etc
	// now some map, where we can get users from
	UserData   map[int]models.User
	LastUserId int
}

func NewDataBase() (*Database, error) {
	// error is possible error from database
	db := new(Database)
	// some code(real code working with db)

	data := make(map[int]models.User)
	db.LastUserId = 50
	h := sha256.New()
	h.Write([]byte("pass"))

	for i := 0; i < db.LastUserId; i++ {
		data[i] = models.User{i, "test_user_" + strconv.Itoa(i), "kek@lol.kl", string(h.Sum(nil)), 0, 1, 0, "files/avatars/" + strconv.Itoa(i) + ".jpg"}
	}
	db.UserData = data
	return db, nil
}
