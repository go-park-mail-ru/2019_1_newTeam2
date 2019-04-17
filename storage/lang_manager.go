package storage

import "github.com/user/2019_1_newTeam2/models"

func (db *Database) GetLangs() ([]models.Language, bool, error) {
	results, err := db.Conn.Query(GetLangs)
	if err != nil {
		return []models.Language{}, false, err
	}

	langs := make([]models.Language, 0)
	for results.Next() {
		lang := models.Language{}
		err = results.Scan(&lang.ID, &lang.Name)
		if err != nil {
			return []models.Language{}, false, nil
		}
		langs = append(langs, lang)
	}
	return langs, true, nil
}
