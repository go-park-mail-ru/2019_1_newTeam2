package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Secret       string   `json:"secret,omitempty"`
	Port         string   `json:"port"`
	AuthHost     string   `json:"authHost"`
	AuthPort     string   `json:"authPort"`
	ScorePort    string   `json:"commonDbServPort,omitempty"`
	UploadPath   string   `json:"uploadPath,omitempty"`
	AvatarsPath  string   `json:"avatarsPath,omitempty"`
	AllowedHosts []string `json:"AllowedHosts,omitempty"`
	ScoreHost    string   `json:"scoreHost"`
	DBName       string   `json:"dbName,omitempty"`
	DBUser       string   `json:"dbUser,omitempty"`
	DBPassUser   string   `json:"dbPassUser,omitempty"`
	DBHost       string   `json:"dbHost,omitempty"`
	RoomSize     int      `json:"roomSize,omitempty"`
	Akid         string   `json:"akid,omitempty"`
	CloudSecret  string   `json:"cloudSecret,omitempty"`
	Bucket       string   `json:"bucket,omitempty"`
}

func NewConfig(pathToConfig string) (*Config, error) {
	conf := new(Config)
	configFile, err := os.Open(pathToConfig)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = configFile.Close()
	}()
	//err = easyjson.UnmarshalFromReader(configFile, conf)
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
