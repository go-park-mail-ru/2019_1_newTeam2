package filesystem_test

import (
	"github.com/user/2019_1_newTeam2/shared/pkg/filesystem"
	"os"
	"testing"
)

type TestCreateDirCase struct {
	path    string
	success bool
}

func TestCreateDir(t *testing.T) {
	cases := []TestCreateDirCase{
		{
			path:    "/tmp/dir/",
			success: true,
		},
		{
			path:    "/etc/dir",
			success: false,
		},
		{
			path:    "../filesystem",
			success: false,
		},
	}
	for _, testCase := range cases {
		err := filesystem.CreateDir(testCase.path)
		if testCase.success {
			if err != nil {
				t.Error(err)
			} else {
				_, err = os.Stat(testCase.path)
				if os.IsNotExist(err) {
					t.Error(err)
				}
			}
		}
	}
}
