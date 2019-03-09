package storage

import "fmt"

type UserStorage struct {
	Data   map[int]User
	LastId int
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

func (users UserStorage) Login(username string, password string) (string, error) {
	fmt.Println(username, password)
	return "lolkek4eburek", nil
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
	password string, langid int, pronounceOn int) (bool, error) {
	id := users.LastId
	fmt.Println(users.LastId)

	users.Data[id] = User{id, username, email, password, langid, pronounceOn, 0}
	return true, nil
}

func (users UserStorage) DeleteUserById(userID int) (bool, error) {
	delete(users.Data, userID)
	return true, nil
}

func (users UserStorage) UpdateUserById(userID int, username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {
	users.Data[userID] = User{userID, username, email, password, langid, pronounceOn, users.Data[userID].Score}
	return true, nil
}

func (users UserStorage) GetAllUser() ([]User, error) {
	allUsers := make([]User, 0)
	for _, i := range users.Data {
		allUsers = append(allUsers, i)
	}
	return allUsers, nil
}
