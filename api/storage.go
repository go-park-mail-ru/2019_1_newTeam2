package main

type userStorage struct {
	data []User
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	langID      int    `json:"langID, int"`
	pronounceON int    `json:"pronounceOn, int"`
}

func (users userStorage) getUserByID(userID int) (User, bool, error) {
	for _, i := range users.data {
		if i.ID == userID {
			return i, true, nil
		}
	}
	return User{}, false, nil
}
