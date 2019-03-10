package storage

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type UserStorage struct {
	Data   map[int]User
	LastId int
}

type User struct {
	ID          int    `json:"id,omitempty"`
	Username    string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	LangID      int    `json:"langID, int"`
	PronounceON int    `json:"pronounceOn, int"`
	Score       int    `json:"score, int"`
	AvatarPath  string `json:"path"`
}

func (users UserStorage) IsLogin(w http.ResponseWriter, r *http.Request, username string, password string) bool {

	return false
}

func (users UserStorage) Login(username string, password string) (string, error) {
	SECRET := []byte("kekusmaxima")
	fmt.Println("DAta is ", username, password)
	for _, i := range users.Data {
		if i.Username == username {
			// h := sha256.New()
			// h.Write([]byte(password))
			// if string(h.Sum(nil)) == i.Password {
			if password == i.Password {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"username": username,
					"password": password,
				})
				str, _ := token.SignedString(SECRET)
				return str, nil
			} else {
				return "", fmt.Errorf("Error bad password")
			}
		}
	}
	return "", fmt.Errorf("Error not user")
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

	users.Data[id] = User{id, username, email, password, langid, pronounceOn, 0, "uploads/avatars/1.jpg"}
	return true, nil
}

func (users UserStorage) DeleteUserById(userID int) (bool, error) {
	delete(users.Data, userID)
	return true, nil
}

func (users UserStorage) UpdateUserById(userID int, username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {
	users.Data[userID] = User{userID, username, email, password, langid, pronounceOn, users.Data[userID].Score, "uploads/avatars/1.jpg"}
	return true, nil
}

func (users UserStorage) GetAllUser() ([]User, error) {
	allUsers := make([]User, 0)
	for _, i := range users.Data {
		allUsers = append(allUsers, i)
	}
	return allUsers, nil
}

func (users UserStorage) AddImage(path string, userID int) error {
	_, ok := users.Data[userID]
	if !ok {
		return fmt.Errorf("no such user")
	}
	user := users.Data[userID]
	user.AvatarPath = path
	fmt.Println(path)
	users.Data[userID] = user
	return nil
}
