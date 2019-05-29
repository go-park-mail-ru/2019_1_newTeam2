package filesystem_test

import (
	"bytes"
	"fmt"
	"github.com/user/2019_1_newTeam2/shared/filesystem"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewSuccessRequest(text string, filename string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.SetBoundary("boundary")
	part, _ := writer.CreateFormFile("file", filename)
	_, _ = part.Write([]byte(text))
	_ = writer.Close()
	req, _ := http.NewRequest("POST", "localhost:8090", body)
	req.Header.Add("Content-Type", "multipart/form-data; charset=utf-8; boundary=\"boundary\"")
	return req
}

func NewFailRequest(text string, filename string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.SetBoundary("boundary")
	part, _ := writer.CreateFormFile("file", filename)
	_, _ = part.Write(nil)
	_ = writer.Close()
	req, _ := http.NewRequest("POST", "localhost:8090", body)
	req.Header.Add("Content-Type", "multipart/form-data; charset=utf-8; boundary=\"boundary\"")
	return req
}

var (
	SuccessCallback = func(header multipart.FileHeader) error {
		return nil
	}
	FailCallback = func(header multipart.FileHeader) error {
		return fmt.Errorf("fail")
	}
)

func TestUploadFileSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	req := NewSuccessRequest("hello", "text.txt")
	path, err := filesystem.UploadFile(w, req, SuccessCallback, "/", "/tmp/")
	if err != nil {
		t.Error(err)
	}

	if path != "files/tmp/text.txt" {
		t.Errorf("file names don't match, Expected: files/text.txt\n Got %v", path)
	}
}

func TestUploadFileWrongReq(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, httptest.DefaultRemoteAddr, nil)
	_, err := filesystem.UploadFile(w, req, SuccessCallback, "/", "/")
	if err == nil {
		t.Error("Unhandled error")
	}
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Error("Wrong status after error")
	}
}

func TestUploadFileCallBack(t *testing.T) {
	w := httptest.NewRecorder()
	req := NewSuccessRequest("hello", "text.txt")
	_, err := filesystem.UploadFile(w, req, FailCallback, "/", "/tmp/")
	if err == nil {
		t.Error("Unhandled error")
	}
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Error("Wrong status after error")
	}
}
