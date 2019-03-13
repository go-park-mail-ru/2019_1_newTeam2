package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/gorilla/mux"

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

func (suite *UserHandlerTestSuite) TestGetUserCorrect() {
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
}

func (suite *UserHandlerTestSuite) TestUsersPaginate() {
	cases := []TestUsersPaginateCase{
		TestUsersPaginateCase{
			t: []models.UserTableElem{
				&models.UserTableElem{
					Username: "vasya",
					Score: 5,
				},
				&models.UserTableElem{
					Username: "petya",
					Score: 7,
				},
				&models.UserTableElem{
					Username: "kolya",
					Score: 0,
				},
				&models.UserTableElem{
					Username: "tanya",
					Score: 9,
				},
			},
			response: "200 OK"
			row: 1,
			page: 5,
		},
		TestUsersPaginateCase{
			
		},
		TestUsersPaginateCase{
			
		},
	}
	/*t := []*models.UserTableElem{
		&models.UserTableElem{
			Username: "vasya",
			Score: 5,
			},
		&models.UserTableElem{
			Username: "petya",
			Score: 7,
			},
		&models.UserTableElem{
			Username: "kolya",
			Score: 0,
			},
		&models.UserTableElem{
			Username: "tanya",
			Score: 9,
			},
	}
	suite.dataBase.EXPECT().GetUsers(1, 5).Return(t, nil)

	vars := map[string]string{
		"rows": "5",
		"page": "1",
	}

	r, _ := http.NewRequest("GET", "/users", nil)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZhc3lhIiwicGFzc3dvcmQiOiIxMjM0NSIsImlkIjoxfQ.CShosAAiK5Dea_7UJ_M2omHyyOtPcmVJkzbiOFWgtn4"
	PlaceTokenToRequest(token, r)
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()
	suite.underTest.UsersPaginate(w, r)
	response := w.Result()

	suite.Equal("200 OK", response.Status)
	defer response.Body.Close()
	result := new(models.UserTableElem)
	json.NewDecoder(response.Body).Decode(result)

	suite.Equal(t, result)*/


}

func (suite *UserHandlerTestSuite) UpdateUser() {

}

func (suite *UserHandlerTestSuite) DeleteUser() {
	
}