package config_test

import (
	"encoding/json"
	"github.com/user/2019_1_newTeam2/shared/pkg/config"
	"log"
	"os"
	"reflect"
	"testing"
)

type ConfigTestCase struct {
	path  string
	ifErr bool
}

func TestNewConfig(t *testing.T) {
	cases := []ConfigTestCase{
		{
			path:  "/dev/null",
			ifErr: true,
		},
		{
			path:  "/etc/gdb",
			ifErr: true,
		},
		{
			path:  "/tmp/config_api.json",
			ifErr: false,
		},
	}
	file, _ := os.OpenFile("/tmp/config_api.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	enc := json.NewEncoder(file)
	cnf := &config.Config{
		Secret: "12345",
		Port:   "9090",
	}
	_ = enc.Encode(cnf)

	for _, it := range cases {
		curConf, err := config.NewConfig(it.path)
		if it.ifErr {
			if err == nil {
				t.Error("should be error")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(cnf, curConf) {
				log.Println(curConf)
				log.Println(cnf)
				t.Error("should be equal configs, but no")
			}
		}
	}
}
