package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/server"
)

type cases struct {
	ExpectedResponse int
	Input            string
}

func TestSignup(t *testing.T) {
	server := server.TestServer()
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
	server := server.TestServer()
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

func TestNegativeCheck(t *testing.T) {
	server := server.TestServer()
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

func TestLogout(t *testing.T) {
	server := server.TestServer()
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

func TestPositiveCheck(t *testing.T) {
	server := server.TestServer()
	var testCase = cases{
		ExpectedResponse: 200,
		Input:            `{"username": "test_user_1", "password": "pass"}`,
	}
	regRequest, regErr := http.NewRequest("GET", "/session/", strings.NewReader(testCase.Input))
	if regErr != nil {
		panic(regErr.Error())
	}
	regRequest.Header.Set("Content-Type", "application/json")

	regRequest.AddCookie(&http.Cookie{Name: "session_id", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RfdXNlcl8xIiwicGFzc3dvcmQiOiJwYXNzIiwiaWQiOiIxIn0.Pbxpesqv_dPYidTNHxB17f5boVg7kWUrQhXH-G8iCTw"})

	TestResponseRecorder := httptest.NewRecorder()
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/session/", server.IsLogin)
	server.Router.ServeHTTP(TestResponseRecorder, regRequest)

	if TestResponseRecorder.Code != testCase.ExpectedResponse {
		t.Errorf("Error in \"positive\" IsLogin: expected %v, have %v!\n", testCase.ExpectedResponse, TestResponseRecorder.Code)
	}
}
