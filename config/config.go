package config

import (
	"fmt"
	"github.com/micro/go-config"
)

type Config struct {
	Secret string	`json:"secret"`
	Port string	`json:"port"`
	UploadPath string `json:"uploadPath"`
	AvatarsPath string `json:"avatarsPath"`
	// anything u want
}

func NewConfig(pathToConfig string) (*Config, error) {
	conf := new(Config)
	err := config.LoadFile(pathToConfig)
	if err != nil {
		return nil, fmt.Errorf("error loading config")
	}
	err = config.Scan(&conf)
	if err != nil {
		return nil, fmt.Errorf("error scanning config")
	}
	fmt.Println(conf)
	return conf, nil
}