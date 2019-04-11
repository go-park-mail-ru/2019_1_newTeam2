package storage

import "github.com/user/2019_1_newTeam2/models"

func (db *Database)GetDicts(userId int, page int, rowsNum int) ([]models.DictReduced, bool, error) {
	dicts := make([]models.DictReduced, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)
	rows, err := db.Conn.Query(DictsPaginate, userId, rowsNum, offset)
	if err != nil {
		return dicts, false, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		i++
		dict := models.DictReduced{}
		err := rows.Scan(&dict.ID, &dict.Name, &dict.Description, &dict.UserId)
		// TODO(sergeychur): say about userId, may be useful, if no delete
		if err != nil{
			return dicts, false, err
		}
		dicts = append(dicts, dict)
	}
	if i == 0 {
		return dicts, false, nil
	}
	return dicts, true, nil
}
func (db *Database) GetDict(dictId int) (models.DictReduced, bool, error) {
	dict := models.DictReduced{}
	row := db.Conn.QueryRow(GetDictById, dictId)
	err := row.Scan(&dict.ID, &dict.Name, &dict.Description, &dict.UserId)
	if err != nil {
		return models.DictReduced{}, false, err
	}
	return dict, true, nil
}
