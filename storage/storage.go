package storage

type UserStorage struct {
	Data map[int]User
}

type User struct {
	ID          int    `json:"id,omitempty"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	LangID      int    `json:"langID, int"`
	PronounceON int    `json:"pronounceOn, int"`
	Score       int    `json:"score, int"`
}

func (users UserStorage) GetUserByID(userID int) (User, bool, error) {
	for _, i := range users.Data {
		if i.ID == userID {
			return i, true, nil
		}
	}
	return User{}, false, nil
}

func (users UserStorage) UserRegistration(username string, email string,
	password string, langid int, pronounceon int) (bool, error) {
	id := len(users.Data)
	users.Data[id] = User{id, username, email, password, langid, pronounceon, 0}
	return true, nil
}

func (users UserStorage) DeleteUserById(userID int) (bool, error) {
	delete(users.Data, userID)
	return true, nil
}

func (users UserStorage) UpdateUserById(userID int, username string, email string,
	password string, langid int, pronounceon int) (bool, error) {
	users.Data[userID] = User{userID, username, email, password, langid, pronounceon, users.Data[userID].Score}
	return true, nil
}

func (users UserStorage) GerAllUser() ([]User, error) {
	allUsers := make([]User, 0)
	for _, i := range users.Data {
		allUsers = append(allUsers, i)
	}
	return allUsers, nil
}
