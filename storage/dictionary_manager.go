package storage

import (
	"database/sql"
	"fmt"

	"github.com/tealeg/xlsx"
	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) DictionaryDelete(DictID int) error {
	rows, err := db.Conn.Query(GetLibraryIDByDictionaryID, DictID)
	if err != nil {
		db.Logger.Log(err)
		return err
	}

	for rows.Next() {
		var card_id int
		err := rows.Scan(&card_id)
		if err != nil {
			db.Logger.Log(err)
			return err
		}
		db.DecrementCount(card_id)
	}

	_, DeleteErr := db.Conn.Exec(
		DeleteDictionary,
		DictID,
	)
	if DeleteErr != nil {
		db.Logger.Log(DeleteErr)
		return fmt.Errorf("DeleteErr: dictionary not deleted")
	}
	return nil
}

func (db *Database) DictionaryUpdate(DictID int, Name string, Description string) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		db.Logger.Log("DictionaryUpdate: transaction error - ", err)
		return fmt.Errorf("DictionaryUpdate: transaction error")
	}
	_, UpdateErr := tx.Exec(
		UpdateDictionary,
		Name,
		Description,
		DictID,
	)
	if UpdateErr != nil {
		db.Logger.Log(UpdateErr)
		tx.Rollback()
		return fmt.Errorf("DictionaryUpdate: user not update")
	}
	tx.Commit()
	return nil
}

func (db *Database) DictionaryCreate(UserID int, Name string, Description string, Cards []models.Card) (models.DictionaryInfoPrivilege, error) {
	tx, err := db.Conn.Begin()

	result, CreateErr := tx.Exec(
		CreateEmptyDictionary,
		Name,
		Description,
		UserID,
	)
	if CreateErr != nil {
		db.Logger.Log(CreateErr)
		tx.Rollback()
		return models.DictionaryInfoPrivilege{}, fmt.Errorf("CreateErr: user not create")
	}
	tx.Commit()
	lastID, GetIDErr := result.LastInsertId()
	if GetIDErr != nil {
		db.Logger.Log(GetIDErr)
		return models.DictionaryInfoPrivilege{}, fmt.Errorf("GetIDErr: can`t get last dict id")
	}
	for _, it := range Cards {
		err := db.SetCardToDictionary(int(lastID), it)
		if err != nil {
			return models.DictionaryInfoPrivilege{}, err
		}
	}
	dict, _, err := db.GetDict(int(lastID))
	if err != nil {
		db.Logger.Log(err)
		return models.DictionaryInfoPrivilege{}, err
	}
	return dict, nil
}

func (db *Database) GetDicts(userId int, page int, rowsNum int) ([]models.DictionaryInfo, bool, error) {
	dicts := make([]models.DictionaryInfo, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)

	rows, err := db.Conn.Query(DictsPaginate, userId, rowsNum, offset)
	if err != nil {
		db.Logger.Log(err)
		return dicts, false, err
	}
	defer rows.Close()
	i := 0

	for rows.Next() {
		i++
		dict := models.DictionaryInfo{}
		err := rows.Scan(&dict.ID, &dict.Name, &dict.Description, &dict.UserId)
		if err != nil {
			db.Logger.Log(err)
			return dicts, false, err
		}
		dicts = append(dicts, dict)
	}

	if i == 0 {
		return dicts, false, nil
	}
	return dicts, true, nil
}

func (db *Database) GetDict(dictId int) (models.DictionaryInfoPrivilege, bool, error) {
	dict := models.DictionaryInfoPrivilege{}
	row := db.Conn.QueryRow(GetDictById, dictId)
	err := row.Scan(&dict.ID, &dict.Name, &dict.Description, &dict.UserId)
	if err != nil {
		return models.DictionaryInfoPrivilege{}, false, err
	}
	return dict, true, nil
}

func (db *Database) BorrowDictById(dictId int, thiefId int) (int, models.DictionaryInfo, error) {
	ownerId := 0
	row := db.Conn.QueryRow(GetDictOwner, dictId)
	err := row.Scan(&ownerId)
	if err == sql.ErrNoRows {
		return 0, models.DictionaryInfo{}, ErrNotFound
	}
	dict := models.DictionaryInfo{}
	row = db.Conn.QueryRow(DictBorrowProc, dictId, thiefId)
	err = row.Scan(&dict.ID, &dict.Name, &dict.Description, &dict.UserId)
	if err != nil {
		return 0, models.DictionaryInfo{}, DBerror
	}
	return ownerId, dict, nil
}

func (db *Database) FillDictionaryFromXLSX(dictId int, pathToFile string) error {
	xlsDict, err := xlsx.OpenFile(pathToFile)
	if err != nil {
		db.Logger.Log("file not found: ", err)
		return err
	}


	var language []string
	for _, cell := range xlsDict.Sheets[0].Rows[0].Cells {
		language = append(language, cell.String())
	}

	if (len(language) != 2) {
		return fmt.Errorf("bad file")
	}

	lang1, err := db.GetLangByName(language[0])
	if err != nil {
		db.Logger.Log("FillDictionaryFromXLSX: language", language[0], "not found")
		return err
	}
	lang2, err := db.GetLangByName(language[1])
	if err != nil {
		db.Logger.Log("FillDictionaryFromXLSX: language", language[1], "not found")
		return err
	}

	for _, row := range xlsDict.Sheets[0].Rows[1:] {
		var data []string
		for _, cell := range row.Cells {
			data = append(data, cell.String())
		}
		if (len(data) != 2) {
			return fmt.Errorf("bad file")
		}
		word1 := models.Word{data[0], lang1.ID}
		word2 := models.Word{data[1], lang2.ID}
		card := models.Card{0, &word1, &word2, 0,}
		err = db.SetCardToDictionary(int(dictId), card)
		if err != nil {
			return err
		}
	}
	return nil
}
