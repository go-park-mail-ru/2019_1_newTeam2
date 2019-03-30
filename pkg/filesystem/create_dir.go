package filesystem

import (
	"os"
)

func CreateDir(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
