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

func UploadFileToCloud(w http.ResponseWriter, r *http.Request, callback func(header multipart.FileHeader) error, svc *s3.S3, path string, bucket string) (string, error) {
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
	grantRead := `public-read`
	key := path + "/" + handle.Filename
	_, err = svc.PutObject(&s3.PutObjectInput{
		ACL:    &grantRead,
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		return "", err
	}
	return key, nil
}
