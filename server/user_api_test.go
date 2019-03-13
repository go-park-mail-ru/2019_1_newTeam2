package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	// "github.com/gorilla/mux"

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

func (suite *UserHandlerTestSuite) TestGetUser() {
	t := &models.User{
		ID : 1,
		Username : "vasya",
		Email : "vasya@gmail.com",
		Password : "12345",
		LangID : 0,
		PronounceON : 0,
		Score : 15,
		AvatarPath : "",
	}
	suite.dataBase.EXPECT().GetUserByID(1).Return(*t, true, nil)
	r, _ := http.NewRequest("GET", "/users/1", nil)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZhc3lhIiwicGFzc3dvcmQiOiIxMjM0NSIsImlkIjoxfQ.CShosAAiK5Dea_7UJ_M2omHyyOtPcmVJkzbiOFWgtn4"
	cookie := &http.Cookie{
			Name:  "session_id",
			Value: token,
		}
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(5 * time.Hour)
		cookie.HttpOnly = false
		cookie.Secure = false
	r.AddCookie(cookie)
	w := httptest.NewRecorder()
	suite.underTest.GetUser(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)
	defer response.Body.Close()
	result := new(models.User)
	json.NewDecoder(response.Body).Decode(result)

	suite.Equal(t, result)
}

func (suite *UserHandlerTestSuite) TestUsersPaginate() {

}

func (suite *UserHandlerTestSuite) UpdateUser() {

}

func (suite *UserHandlerTestSuite) DeleteUser() {
	
}