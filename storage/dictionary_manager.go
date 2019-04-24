package storage

import (
	"database/sql"
	"fmt"

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
	_, UpdateErr := db.Conn.Exec(
		UpdateDictionary,
		Name,
		Description,
		DictID,
	)
	if UpdateErr != nil {
		db.Logger.Log(UpdateErr)
		return fmt.Errorf("UpdateErr: user not update")
	}

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
