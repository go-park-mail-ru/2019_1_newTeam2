package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"bytes"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/user/2019_1_newTeam2/mock_database"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/server"
	"github.com/user/2019_1_newTeam2/config"
)

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

type UserHandlerTestSuite struct {
	suite.Suite
	dataBase *mock_database.MockDBInterface
	underTest     *server.Server
}

func (suite *UserHandlerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.dataBase = mock_database.NewMockDBInterface(mockCtrl)

	server := new(server.Server)
	config := new(config.Config)
	config.Secret = "12345"
	config.Port = ":8090"
	config.UploadPath = "/temp"
	config.AvatarsPath = ""
	server.ServerConfig = config
	server.DB = suite.dataBase
	suite.underTest = server
}

func PlaceTokenToRequest(token string, r *http.Request) {
	cookie := &http.Cookie{
			Name:  "session_id",
			Value: token,
		}
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(5 * time.Hour)
		cookie.HttpOnly = false
		cookie.Secure = false
	r.AddCookie(cookie)
}

type TestGetUserCase struct {
	t models.User
	response string
	id int
	exists bool
	err error
}

type TestErr struct {
	str string
}

func (err *TestErr) Error() string {
	return err.str
}

func (suite *UserHandlerTestSuite) TestGetUser() {
	cases := []TestGetUserCase{
		TestGetUserCase{
			t: models.User{
				ID: 1,
				Username: "vasya",
				Email: "vasya@gmail.com",
				Password: "12345",
				LangID: 0,
				PronounceON: 0,
				Score: 15,
				AvatarPath: "",
			},
			response: "200 OK",
			id: 1,
			err: nil,
			exists: true,
		},
		TestGetUserCase{
			t: models.User{},
			response: "404 Not Found",
			id: 1,
			err: nil,
			exists: false,
		},
		TestGetUserCase{
			t: models.User{},
			response: "500 Internal Server Error",
			id: 1,
			err: &TestErr{
				str: "db error",
				},
			exists: false,
		},
	}

	for _, item := range(cases) {
		suite.dataBase.EXPECT().GetUserByID(item.id).Return(item.t, item.exists, item.err)
		r, _ := http.NewRequest("GET", "/users/", nil)
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZhc3lhIiwicGFzc3dvcmQiOiIxMjM0NSIsImlkIjoxfQ.CShosAAiK5Dea_7UJ_M2omHyyOtPcmVJkzbiOFWgtn4"
		PlaceTokenToRequest(token, r)
		w := httptest.NewRecorder()
		suite.underTest.GetUser(w, r)

		response := w.Result()
		suite.Equal(item.response, response.Status)
		if item.exists &&  item.err == nil {
			defer response.Body.Close()
			result := new(models.User)
			json.NewDecoder(response.Body).Decode(result)

			suite.Equal(item.t, *result)
		}
	}
}

type TestUsersPaginateCase struct {
	t []models.UserTableElem
	response string
	err error
	row int
	page int
	strRow string
	strPage string
	rowsURL map[string]string
	pageURL map[string]string
}



func (suite *UserHandlerTestSuite) TestUsersPaginate() {
	cases := []TestUsersPaginateCase{
		TestUsersPaginateCase{
			t: []models.UserTableElem{
				models.UserTableElem{
					Username: "vasya",
					Score: 5,
				},
				models.UserTableElem{
					Username: "petya",
					Score: 7,
				},
				models.UserTableElem{
					Username: "kolya",
					Score: 0,
				},
				models.UserTableElem{
					Username: "tanya",
					Score: 9,
				},
			},
			row: 1,
			page: 5,
			strRow: "1",
			strPage: "5",
			response: "200 OK",
			err: nil,

		},

		TestUsersPaginateCase{
			t: []models.UserTableElem{
				models.UserTableElem{
					Username: "vasya",
					Score: 5,
				},
				models.UserTableElem{
					Username: "petya",
					Score: 7,
				},
				models.UserTableElem{
					Username: "kolya",
					Score: 0,
				},
				models.UserTableElem{
					Username: "tanya",
					Score: 9,
				},
			},
			row: 1,
			strRow: "",
			page: 5,
			strPage: "",
			response: "400 Bad Request",
			err: &TestErr{
				str: "no query",
				},

		},

		TestUsersPaginateCase{
			t: []models.UserTableElem{
				models.UserTableElem{
					Username: "vasya",
					Score: 5,
				},
				models.UserTableElem{
					Username: "petya",
					Score: 7,
				},
				models.UserTableElem{
					Username: "kolya",
					Score: 0,
				},
				models.UserTableElem{
					Username: "tanya",
					Score: 9,
				},
			},
			row: 1,
			strRow: "ede",
			page: 5,
			strPage: "ede",
			response: "400 Bad Request",
			err: &TestErr{
				str: "bad query",
				},

		},

		TestUsersPaginateCase{
			t: []models.UserTableElem{
				models.UserTableElem{
					Username: "vasya",
					Score: 5,
				},
				models.UserTableElem{
					Username: "petya",
					Score: 7,
				},
				models.UserTableElem{
					Username: "kolya",
					Score: 0,
				},
				models.UserTableElem{
					Username: "tanya",
					Score: 9,
				},
			},
			row: 1,
			strRow: "1",
			page: 5,
			strPage: "5",
			response: "404 Not Found",
			err: &TestErr{
				str: "not found",
				},

		},
	}

	for _, item := range(cases) {
		suite.dataBase.EXPECT().GetUsers(item.page, item.row).Return(item.t, item.err)
		r, _ := http.NewRequest("GET", "/users", nil)
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZhc3lhIiwicGFzc3dvcmQiOiIxMjM0NSIsImlkIjoxfQ.CShosAAiK5Dea_7UJ_M2omHyyOtPcmVJkzbiOFWgtn4"
		PlaceTokenToRequest(token, r)
		q := r.URL.Query()
		q.Add("rows", item.strRow)
		q.Add("page", item.strPage)
		r.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		suite.underTest.UsersPaginate(w, r)
		response := w.Result()
		suite.Equal(item.response, response.Status)
		if item.err == nil {
			defer response.Body.Close()
			result := []models.UserTableElem{}

			json.NewDecoder(response.Body).Decode(&result)
			suite.Equal(item.t, result)
		}
	}
}

type TestUpdateUserCase struct {
	t models.User
	response string
	exists bool
	err error
}

func (suite *UserHandlerTestSuite) UpdateUser() {
	cases := []TestUpdateUserCase{
		TestUpdateUserCase{
			t: models.User{
				ID: 1,
				Username: "vasya",
				Email: "vasya@gmail.com",
				Password: "12345",
				LangID: 0,
				PronounceON: 0,
				Score: 15,
				AvatarPath: "",
			},
			response: "200 OK",
			err: nil,
			exists: true,
		},
		TestUpdateUserCase{
			t: models.User{},
			response: "404 Not Found",
			err: nil,
			exists: false,
		},
		TestUpdateUserCase{
			t: models.User{},
			response: "500 Internal Server Error",
			err: &TestErr{
				str: "db error",
				},
			exists: false,
		},
	}

	for _, item := range(cases) {
		suite.dataBase.EXPECT().GetUserByID(item.t.ID).Return(item.t, item.exists, item.err)
		suite.dataBase.EXPECT().UpdateUserById(item.t.ID, item.t.Username, item.t.Email, 
			item.t.Password, item.t.LangID, item.t.PronounceON)
		body, _ := json.Marshal(item.t)
		r, _ := http.NewRequest("PATCH", "/users/", bytes.NewBuffer(body))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZhc3lhIiwicGFzc3dvcmQiOiIxMjM0NSIsImlkIjoxfQ.CShosAAiK5Dea_7UJ_M2omHyyOtPcmVJkzbiOFWgtn4"
		PlaceTokenToRequest(token, r)

		w := httptest.NewRecorder()

		suite.underTest.UpdateUser(w, r)

		response := w.Result()
		suite.Equal(item.response, response.Status)
	}
}

type TestDeleteUserCase struct {
	t models.User
	id int
	response string
	exists bool
	err error
}

func (suite *UserHandlerTestSuite) DeleteUser() {
	cases := []TestDeleteUserCase{
		TestDeleteUserCase{
			t: models.User{
				ID: 1,
				Username: "vasya",
				Email: "vasya@gmail.com",
				Password: "12345",
				LangID: 0,
				PronounceON: 0,
				Score: 15,
				AvatarPath: "",
			},
			response: "200 OK",
			id: 1,
			err: nil,
			exists: true,
		},
		TestDeleteUserCase{
			t: models.User{},
			response: "404 Not Found",
			id: 1,
			err: nil,
			exists: false,
		},
		TestDeleteUserCase{
			t: models.User{},
			response: "500 Internal Server Error",
			id: 1,
			err: &TestErr{
				str: "db error",
				},
			exists: false,
		},

	}

	for _, item := range(cases) {
		suite.dataBase.EXPECT().GetUserByID(item.id).Return(item.t, item.exists, item.err)
		suite.dataBase.EXPECT().UpdateUserById(item.t.ID, item.t.Username, item.t.Email, 
			item.t.Password, item.t.LangID, item.t.PronounceON)
		body, _ := json.Marshal(item.t)
		r, _ := http.NewRequest("PATCH", "/users/", bytes.NewBuffer(body))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZhc3lhIiwicGFzc3dvcmQiOiIxMjM0NSIsImlkIjoxfQ.CShosAAiK5Dea_7UJ_M2omHyyOtPcmVJkzbiOFWgtn4"
		PlaceTokenToRequest(token, r)

		w := httptest.NewRecorder()

		suite.underTest.UpdateUser(w, r)

		response := w.Result()
		suite.Equal(item.response, response.Status)
	}
}