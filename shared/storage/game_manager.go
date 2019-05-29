package storage

import "github.com/user/2019_1_newTeam2/shared/models"

func (db *Database) UpdateFrequencies(results models.GameResults) (error, bool) {
	allFound := true
	tx, err := db.Conn.Begin()
	if err != nil {
		return err, false
	}
	defer tx.Rollback()
	stmt, err := db.Conn.Prepare(UpdateFrequency)
	defer stmt.Close()
	if err != nil {
		return err, false
	}
	for _, update := range results {
		guessed := 0
		if update.Correct {
			guessed++
		}
		res, err := stmt.Exec(guessed, update.ID)
		if err != nil {
			return err, false
		}
		num, err := res.RowsAffected()
		if num == 0 {
			allFound = false
		}
	}
	err = tx.Commit()
	if err != nil {
		return err, false
	}
	return nil, allFound
}
