package storage

import "github.com/user/2019_1_newTeam2/shared/models"

func (db *Database) GetLangs() (models.Langs, bool, error) {
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

func (db *Database) GetLangByName(LangName string) (models.Language, error) {
	lang := models.Language{}
	result := db.Conn.QueryRow(GetLangById, LangName)
	err := result.Scan(&lang.ID, &lang.Name)
	if err != nil {
		db.Logger.Log("GetLangById: ", err)
		return models.Language{}, err
	}
	return lang, nil
}
