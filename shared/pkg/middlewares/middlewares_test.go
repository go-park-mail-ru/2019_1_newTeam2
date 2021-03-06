package middlewares_test

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/user/2019_1_newTeam2/shared/pkg/middlewares"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

type CorsGetToBody struct {
	uri         string
	method      string
	checkInside bool
	hostInHosts bool
}

func TestCreateCorsMiddleware(t *testing.T) {
	cases := []CorsGetToBody{
		{
			uri:         "http://localhost/uri",
			method:      http.MethodGet,
			checkInside: true,
			hostInHosts: true,
		},
		{
			uri:         "http://localhost/1",
			method:      http.MethodOptions,
			checkInside: false,
			hostInHosts: true,
		},
	}
	hosts := []string{"http://localhost", "https://newwordtrainer.ru"}
	function := middlewares.CreateCorsMiddleware(hosts)

	for _, testCase := range cases {
		here := false
		r := httptest.NewRequest(testCase.method, testCase.uri, nil)
		w := httptest.NewRecorder()
		h := http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				here = true
				w.WriteHeader(http.StatusOK)
			})
		function.Middleware(h).ServeHTTP(w, r)
		if testCase.checkInside {
			if !here {
				t.Error("Failed cors check")
			}
		} else if testCase.hostInHosts && w.Result().StatusCode == http.StatusMethodNotAllowed {
			t.Errorf("Options don't work")
		}
	}
}

func TestCreatePanicRecoveryMiddleware(t *testing.T) {
	function := middlewares.CreatePanicRecoveryMiddleware()
	h := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			panic("panic")
		})
	r := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	w := httptest.NewRecorder()
	caught := true
	defer func() {
		caught = false
		_ = recover()
	}()
	function(h).ServeHTTP(w, r)
	if !caught {
		t.Error("Middleware failed")
	}
	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Error("wrong status code for panic")
	}

}

func PlaceTokenToRequest(token string, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: token,
	}
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(5 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = false
	r.AddCookie(cookie)
}

const correctToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InZhc3lhIiwicGFzc3dvcmQiOiIxMjM0NSIsImlkIjoxfQ.CShosAAiK5Dea_7UJ_M2omHyyOtPcmVJkzbiOFWgtn4"

func GetIdFromCookie(in string, secret []byte) (int, error) {
	token, err := jwt.Parse(in, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["id"].(float64)), nil
	}
	return 0, fmt.Errorf("token invalid")
}

func IsLogined(r *http.Request, secret []byte, cookieField string) bool {
	cookie, err := r.Cookie(cookieField)
	if err != nil {
		return false
	}
	_, err = GetIdFromCookie(cookie.Value, secret)
	return err == nil
}

func TestCreateCheckAuthMiddleware(t *testing.T) {
	ifPlaceToken := []bool{true, false}

	checkFunc := IsLogined
	function := middlewares.CreateCheckAuthMiddleware([]byte("12345"), "session_id", checkFunc)
	inside := false
	h := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			inside = true
		})
	for _, it := range ifPlaceToken {
		r := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
		if it {
			PlaceTokenToRequest(correctToken, r)
		}
		w := httptest.NewRecorder()
		function.Middleware(h).ServeHTTP(w, r)
		if !inside && it {
			t.Error("should have authed")
		}
	}
}

func TestCreateLoggingMiddleware(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	function := middlewares.CreateLoggingMiddleware(buf, "LOG_TEST")
	h := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	r := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	w := httptest.NewRecorder()
	function.Middleware(h).ServeHTTP(w, r)
	str := buf.String()
	log.SetOutput(os.Stderr)
	log.Println(str)
	if !strings.Contains(str, "[LOG_TEST] GET") ||
		!strings.Contains(str, "200 OK") {
		t.Error("Wrong log format")
	}
}
