package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/user/2019_1_newTeam2/pkg/logger"
)

type Database struct {
	// should be sqlx.db etc
	// now some map, where we can get users from
	Conn   *sql.DB
	Logger logger.LoggerInterface
}

func NewDataBase() (*Database, error) {
	// error is possible error from database
	db := new(Database)
	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	db.Logger = logger

	dsn := "root_use" + ":Abc123456*" + "@tcp(localhost:3306)/"
	database, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	err = database.Ping()
	if err != nil {
		log.Println("lol")
	}
	err = createTable(database)
	if err != nil {
		return nil, fmt.Errorf("mysql: could not create database: %v", err)
	}

	db.Conn = database
	return db, nil
}
