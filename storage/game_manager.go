package storage

import "github.com/user/2019_1_newTeam2/models"

func (db *Database) UpdateFrequencies (results models.GameResults) (error, bool) {
	/*allFound := true
	stmt, err := db.Conn.Prepare(UpdateFrequency)
	if err != nil {
		return err, false
	}
	for _,update := range results {

		res, err := stmt.Exec(UpdateFrequency, )
	}*/
	return nil, true
}
