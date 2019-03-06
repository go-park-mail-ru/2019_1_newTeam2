package storage

type UserStorage struct {
	Data []User
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	LangID      int    `json:"langID, int"`
	PronounceON int    `json:"pronounceOn, int"`
}

func (users UserStorage) GetUserByID(userID int) (User, bool, error) {
	for _, i := range users.Data {
		if i.ID == userID {
			return i, true, nil
		}
	}
	return User{}, false, nil
}
