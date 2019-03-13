package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Secret      string `json:"secret"`
	Port        string `json:"port"`
	UploadPath  string `json:"uploadPath"`
	AvatarsPath string `json:"avatarsPath"`
	// anything u want
}

func NewConfig(pathToConfig string) (*Config, error) {
	conf := new(Config)
	configFile, err := os.Open(pathToConfig)
	if err != nil {
		return nil, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&conf)
	if err != nil {
		return nil, err
	}
  
	fmt.Println(conf.Port)
	return conf, nil
}
