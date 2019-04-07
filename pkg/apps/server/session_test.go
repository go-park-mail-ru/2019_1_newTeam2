package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type cases struct {
	ExpectedResponse int
	Input            string
}

func TestOptionsSignup(t *testing.T) {
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
	server := TestServer()
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
