package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ircConfig struct {
	Server   string
	Password string
	Port     float64
	Nick     string
	Channels []string
}

type config struct {
	IRC        *ircConfig
	TimeWindow float64
}

var Config *config = new(config)

func loadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, Config)
	return err
}
