package application

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	FilePath        string `json:"filePath"`
	FieldLat        string `json:"fieldLat"`
	FieldLong       string `json:"fieldLong"`
	FieldTime       string `json:"fieldTime"`
	FieldTimeFormat string `json:"fieldTimeFormat"`
}

func LoadConfig(configPath string) (*Config, error) {
	fileContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
