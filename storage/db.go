package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/user/2019_1_newTeam2/pkg/logger"
)

var (
	ErrNotFound error = errors.New("not found")
	DBerror error = errors.New("some db error")
)

type Database struct {
	// should be sqlx.db etc
	// now some map, where we can get users from
	Conn   *sql.DB
	Logger logger.LoggerInterface
}

func NewDataBase(username string, pass string) (*Database, error) {
	// error is possible error from database
	db := new(Database)
	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	db.Logger = logger

	dsn := username + ":" + pass + "@tcp(localhost:3306)/"
	database, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	err = database.Ping()
	// _, err = database.Exec(UseDB)
	// if err != nil {
	// 	return nil, fmt.Errorf("mysql: could not choose db: %v", err)
	// }

	db.Conn = database
	return db, nil
}
