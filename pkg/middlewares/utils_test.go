package middlewares_test

import (
	"bufio"
	"github.com/user/2019_1_newTeam2/pkg/middlewares"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

type recorderMock struct {
}

func (r *recorderMock) Header() http.Header {
	return nil
}

func (r *recorderMock) Write(b []byte) (int, error) {
	return 0, nil
}

func (r *recorderMock) WriteHeader(statusCode int) {

}

func (r *recorderMock) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}

func TestLogResWriterHijack(t *testing.T) {
	w := &recorderMock{}
	lw := middlewares.NewLogResWriter(w)
	_, _, err := lw.Hijack()
	if err != nil {
		t.Error("Oops, error occured, but shouldn't, hijack exists")
	}
}

func TestLogResWriterHijackFail(t *testing.T) {
	w := httptest.NewRecorder()
	lw := middlewares.NewLogResWriter(w)
	_, _, err := lw.Hijack()
	if err == nil {
		t.Error("Should be error")
	}
}
