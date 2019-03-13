package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/user/2019_1_newTeam2/logger"
	"github.com/user/2019_1_newTeam2/models"
)

type Database struct {
	// should be sqlx.db etc
	// now some map, where we can get users from
	UserData   map[int]models.User
	LastUserId int
	Logger     logger.LoggerInterface
}

func NewDataBase() (*Database, error) {
	// error is possible error from database
	db := new(Database)
	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	db.Logger = logger
	// some code(real code working with db)

	data := make(map[int]models.User)
	db.LastUserId = 50

	password := "pass"
	_, err := HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hash error")
	}

	for i := 0; i < db.LastUserId; i++ {
		data[i] = models.User{i, "test_user_" + strconv.Itoa(i), "kek@lol.kl", password, 0, 1, 0, "files/avatars/" + strconv.Itoa(i) + ".jpg"}
	}
	db.UserData = data
	return db, nil
}
