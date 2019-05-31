package filesystem

import (
	"crypto/sha256"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
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

func UploadFileToCloud(w http.ResponseWriter, r *http.Request, callback func(header multipart.FileHeader) error, svc *s3.S3, path string, bucket string, prevKey string) (string, error) {
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

	h := sha256.New()

	_, err = io.Copy(h, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	ext := filepath.Ext(handle.Filename)
	key := fmt.Sprintf("%s%x%s", path, h.Sum(nil), ext)
	grantRead := `public-read`

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

	if prevKey != key {
		_, _ = svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(prevKey),
		})
	}

	return key, nil
}
