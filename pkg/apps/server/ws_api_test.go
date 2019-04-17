package server_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/user/2019_1_newTeam2/mocks/mock_hub"
	"github.com/user/2019_1_newTeam2/pkg/apps/server"
	"github.com/user/2019_1_newTeam2/pkg/config"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWSHandlerSuite(t *testing.T) {
	suite.Run(t, new(WSHandlerTestSuite))
}

type WSHandlerTestSuite struct {
	suite.Suite
	wsHub  *mock_hub.MockIWSCommunicator
	underTest *server.Server
}

func (suite *WSHandlerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()
	suite.wsHub = mock_hub.NewMockIWSCommunicator(mockCtrl)

	server := new(server.Server)
	config := new(config.Config)
	config.Secret = "12345"
	config.Port = ":8090"
	config.UploadPath = "/temp"
	config.AvatarsPath = ""
	server.ServerConfig = config
	server.DB = nil
	server.Hub = suite.wsHub
	suite.underTest = server

	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("TESTLOG: ")
	server.Logger = logger
}

type WSSubscribeTestCase struct {
	responseCode int
	id       int
	err      error
	token    string
	correctToken bool
}

func (suite *WSHandlerTestSuite) TestWSSubscribe() {
	cases := []WSSubscribeTestCase{
		{
			responseCode: http.StatusOK,
			id: 1,
			token: correctToken,
			err : nil,
			correctToken: true,
		},
		{
			responseCode: http.StatusUnauthorized,
			id: 2,
			token: "sss",
			//err: &TestErr{str:"error mes"},
			correctToken: false,
		},
		{
			responseCode: http.StatusBadRequest,
			id: 1,
			token: correctToken,
			err: &TestErr{str:"error mes"},
			correctToken: true,
		},
	}
	for _, testCase := range cases {
		if testCase.correctToken {
			r, _ := http.NewRequest(http.MethodGet, "/subscribe", nil)
			w := httptest.NewRecorder()
			PlaceTokenToRequest(testCase.token, r)
			if testCase.correctToken {
				suite.wsHub.EXPECT().AddClient(w, r, testCase.id).Return(testCase.err)
			}
			suite.underTest.WSSubscribe(w, r)
			response := w.Result()
			suite.Equal(testCase.responseCode, response.StatusCode)
		}
	}
}
