package application

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	FieldUserId            string  `json:"fieldUserId"`
	FilePath               string  `json:"filePath"`
	FieldLat               string  `json:"fieldLat"`
	FieldLong              string  `json:"fieldLong"`
	FieldTime              string  `json:"fieldTime"`
	FieldTimeFormat        string  `json:"fieldTimeFormat"`
	AcceptableDisplacement float64 `json:"acceptableDisplacement"`
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
