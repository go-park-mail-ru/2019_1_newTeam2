package server

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/config"
	"github.com/user/2019_1_newTeam2/models"
)

type TestDatabase struct {
	UserData   map[int]models.User
	LastUserId int
}

func TestServer() *Server {
	server := new(Server)
	server.ServerConfig = &config.Config{
		Secret:      "kekusmaxima",
		Port:        "8090",
		AvatarsPath: "/avatars/",
		UploadPath:  "./upload/",
	}
	newDB, _ := InitTestDataBase()
	server.DB = newDB
	server.Router = mux.NewRouter()
	return server
}

func InitTestDataBase() (*TestDatabase, error) {
	db := new(TestDatabase)
	data := make(map[int]models.User)
	db.LastUserId = 10
	h := sha256.New()
	h.Write([]byte("pass"))
	for i := 0; i < db.LastUserId; i++ {
		data[i] = models.User{i, "test_user_" + strconv.Itoa(i), "kek@lol.kl", string(h.Sum(nil)), 0, 1, 0, "files/avatars/" + strconv.Itoa(i) + ".jpg"}
	}
	db.UserData = data
	return db, nil
}

func (db *TestDatabase) Login(username string, password string, secret []byte) (string, string, error) {
	for _, i := range db.UserData {
		if i.Username == username {
			h := sha256.New()
			h.Write([]byte(password))
			if string(h.Sum(nil)) == i.Password {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"username": username,
					"id":       int64(i.ID),
				})
				str, _ := token.SignedString(secret)
				return str, strconv.Itoa(i.ID), nil
			} else {
				return "", "", fmt.Errorf("Error bad password")
			}
		}
	}
	return "", "", fmt.Errorf("Error not user")
}

func (db *TestDatabase) UserRegistration(username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {
	for _, i := range db.UserData {
		if i.Username == username {
			return false, fmt.Errorf("already reg")
		}
	}
	id := db.LastUserId
	fmt.Println(db.LastUserId)
	h := sha256.New()
	h.Write([]byte(password))
	db.UserData[id] = models.User{id, username, email, string(h.Sum(nil)), langid, pronounceOn, 0, "uploads/avatars/1.jpg"}
	return true, nil
}

func (db *TestDatabase) GetUserByID(userID int) (models.User, bool, error) {
	return models.User{}, true, nil
}
func (db *TestDatabase) DeleteUserById(userID int) (bool, error) {
	return true, nil
}
func (db *TestDatabase) GetUsers(page int, rowsNum int) ([]models.UserTableElem, error) {
	return []models.UserTableElem{}, nil
}
func (db *TestDatabase) AddImage(path string, userID int) error {
	return nil
}
func (db *TestDatabase) UpdateUserById(userID int, username string, email string, password string, langid int, pronounceOn int) (bool, error) {
	return true, nil
}
func (db *TestDatabase) IncUserLastID() {
	return
}
