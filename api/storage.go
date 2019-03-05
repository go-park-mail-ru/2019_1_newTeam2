package main

type userStorage struct {
	data []User
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	languageID  int    `json:"languageId"`
	pronounceON bool   `json:"pronounceOn"`
}

func (users userStorage) getUserByID(userID int) (User, bool, error) {
	var id, languageID int
	var username, email, password string
	var pronounceOn bool

	find := false

	for _, i := range users.data {
		if i.ID == userID {
			find = true
			break
		}
	}

	if find {
		result := User{
			ID:          id,
			Username:    username,
			Email:       email,
			Password:    password,
			languageID:  languageID,
			pronounceON: pronounceOn,
		}
		return result, find, nil
	}
	return User{}, false, nil
}
