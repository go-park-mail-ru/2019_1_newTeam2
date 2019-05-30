package filesystem

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func saveFile(file multipart.File, handle *multipart.FileHeader, path string) (string, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(filepath.Join(path, handle.Filename), data, 0644)
	if err != nil {
		return "", err
	}
	return handle.Filename, nil
}

func UploadFile(w http.ResponseWriter, r *http.Request, callback func(header multipart.FileHeader) error, basePath string, path string) (string, error) {
	file, handle, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}
	err = callback(*handle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}

	defer func() {
		_ = file.Close()
	}()

	retPath, err := saveFile(file, handle, filepath.Join(basePath, path))
	if err != nil {
		return "", err
	}
	return filepath.Join("files/", path, retPath), nil
}

func UploadFileToCloud(w http.ResponseWriter, r *http.Request, callback func(header multipart.FileHeader) error, svc *s3.S3) (string, error) {
	file, handle, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}
	err = callback(*handle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("newteam2backs3backet"),
		Key:    aws.String(handle.Filename),
		Body:   file,
	})
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		return "", err
	}
	return handle.Filename, nil
}
