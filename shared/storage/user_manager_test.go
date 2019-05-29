package storage_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/user/2019_1_newTeam2/shared/storagestorage"
)

func TestSuccessCheckUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	row := sqlmock.NewRows([]string{"ID", "Username", "Email", "Password", "LangID", "PronounceON", "Score", "AvatarPath"}).
		AddRow(1, "Kek", "kek@kek.ru", "kek", 1, 1, 0, "kek")

	mock.ExpectQuery(storage.GetUserByUsernameQuery).WithArgs("kek").WillReturnRows(row)
}

func TestNotFoundUserCheckUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	row := sqlmock.NewRows([]string{"ID", "Username", "Email", "Password", "LangID", "PronounceON", "Score", "AvatarPath"})

	mock.ExpectQuery(storage.GetUserByUsernameQuery).WithArgs("kek").WillReturnRows(row)
}
