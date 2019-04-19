package server_test

import (
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/apps/server"
	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)


type TestDatabase struct {
	UserData   map[int]models.User
	LastUserId int
}

func InitTestServer() *server.Server {
	server := new(server.Server)
	server.ServerConfig = &config.Config{
		Secret:      "kekusmaxima",
		Port:        "8090",
		AvatarsPath: "/avatars/",
		UploadPath:  "./upload/",
	}
	newDB, _ := InitTestDataBase()
	server.DB = newDB
	server.Router = mux.NewRouter()
	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("TESTLOG: ")
	server.Logger = logger
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
func (db *TestDatabase) GetUsers(page int, rowsNum int) ([]models.UserTableElem, bool, error) {
	return []models.UserTableElem{}, true, nil
}
func (db *TestDatabase) AddImage(path string, userID int) error {
	return nil
}
func (db *TestDatabase) UpdateUserById(userID int, username string, email string, langid int, pronounceOn int) (bool, error) {
	return true, nil
}
func (db *TestDatabase) IncUserLastID() {
	return
}

func (db *TestDatabase) GetLangs() ([]models.Language, bool, error) {
	return []models.Language{}, true, nil
}

func (db *TestDatabase) GetCards(userId int, page int, rowsNum int) ([]models.Card, bool, error) {
	cards := make([]models.Card, 0)
	return cards, true, nil
}
func (db *TestDatabase) GetCard(cardId int) (models.Card, bool, error) {
	return models.Card{}, true, nil
}

func (db *TestDatabase) GetDict(dictId int) (models.DictionaryInfoPrivilege, bool, error) {
	return models.DictionaryInfoPrivilege{}, true, nil
}

func (db *TestDatabase) GetDicts(userId int, page int, rowsNum int) ([]models.DictionaryInfo, bool, error) {
	a := make([]models.DictionaryInfo, 0)
	return a, true, nil
}

func (db *TestDatabase) SetCardToDictionary(dictID int, card models.Card) error {
	return nil
}

func (db *TestDatabase) DictionaryDelete(DictID int) error {
	return nil
}

func (db *TestDatabase) DeleteCardInDictionary(cardID int, dictionaryID int) error {
	return nil
}

func (db *TestDatabase) DictionaryCreate(UserID int, Name string, Description string, Cards []models.Card) (models.DictionaryInfoPrivilege, error) {
	return models.DictionaryInfoPrivilege{}, nil
}

func (db *TestDatabase) DictionaryUpdate(DictID int, Name string, Description string) error {
	return nil
}

func (db *TestDatabase) BorrowDictById(dictId int, thiefId int) (int, models.DictionaryInfo, error) {
	return 0, models.DictionaryInfo{}, nil
}

func (db *TestDatabase) GetCardsForGame(dictId int, cardsNum int) ([]models.GameWord, bool, error) {
	return nil, false, nil
}

func (db *TestDatabase) UpdateFrequencies(results models.GameResults) (error, bool) {
	return nil, false
}


type cases struct {
	ExpectedResponse int
	Input            string
}

func TestOptionsSignup(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_101", "email": "email@mail.ru", "password": "pass", "langID": 1, "pronounceOn": 1}`,
	}
	regRequest, regErr := http.NewRequest("OPTIONS", "/users/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/users/", server.SignUpAPI)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestOptionsLogin(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("OPTIONS", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.LoginAPI)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestAlreadySignup(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 400,
		Input:            `{"username": "test_user_5", "email": "email@mail.ru", "password": "pass", "langID": 1, "pronounceOn": 1}`,
	}
	regRequest, regErr := http.NewRequest("POST", "/users/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/users/", server.SignUpAPI)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestSignup(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_101", "email": "email@mail.ru", "password": "pass", "langID": 1, "pronounceOn": 1}`,
	}
	regRequest, regErr := http.NewRequest("POST", "/users/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/users/", server.SignUpAPI)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestLogin(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("POST", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.LoginAPI)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestBadLogin(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 401,
		Input:            `{"username": "test_user_1", "password": "kekpass"}`,
	}
	regRequest, regErr := http.NewRequest("POST", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.LoginAPI)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestNegativeCheck(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 204,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("GET", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.IsLogin)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error in \"positive\" IsLogin: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestOptionsLogout(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("OPTIONS", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.Logout)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestLogout(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("PATCH", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.Logout)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestBadJWTCheck(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 204,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("GET", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	regRequest.AddCookie(&http.Cookie{Name: "session_id", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicGFzc3dvcmQiOiJwYXNzIiwidXNlcm5hbWUiOiJ0ZXN0X3VzZXJfMSJ9.Zi01imDsLhyt6NU03BgACX7v2fyiccQsUd1NvdsffgU"})

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.IsLogin)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error in \"positive\" IsLogin: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestPositiveCheck(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("GET", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	regRequest.AddCookie(&http.Cookie{Name: "session_id", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicGFzc3dvcmQiOiJwYXNzIiwidXNlcm5hbWUiOiJ0ZXN0X3VzZXJfMSJ9.Zi01imDsLhyt6NU03BgACX7v2fyiccQsUd1Nv1UB3cU"})

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.IsLogin)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error in \"positive\" IsLogin: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}

func TestOptionsCheck(t *testing.T) {
	server := InitTestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("OPTIONS", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	regRequest.AddCookie(&http.Cookie{Name: "session_id", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicGFzc3dvcmQiOiJwYXNzIiwidXNlcm5hbWUiOiJ0ZXN0X3VzZXJfMSJ9.Zi01imDsLhyt6NU03BgACX7v2fyiccQsUd1Nv1UB3cU"})

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.IsLogin)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error in \"positive\" IsLogin: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}
