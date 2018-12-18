package main

import (
	"testing"
	"encoding/json"
	"log-detection/application"
	"io/ioutil"
	"os"
)

func Test(t *testing.T) {
	configBytes, err := json.Marshal(application.Config{
		FieldTime:              "time",
		FieldLat:               "lat",
		FieldLong:              "long",
		FilePath:               "./log.log",
		FieldTimeFormat:        "2006-01-02 15:04:05",
		AcceptableDisplacement: 200,
	})
	if err != nil {
		return
	}
	ioutil.WriteFile("./config.json", configBytes, os.ModePerm)
}
