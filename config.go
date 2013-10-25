package config

import (
	"encoding/json"
	"github.com/snormore/gologger"
	"io/ioutil"
)

type Config interface{}

func Register(name string, c Config) error {
	return nil
}

type Configurable struct {
	Config interface{}
}

func Read(filePath string, conf *Config) error {
	logger.Info("Loading configuration from %s...", filePath)
	configJson, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return ReadJson(configJson, conf)
}

func ReadJson(configJson []byte, conf *Config) error {
	return json.Unmarshal(configJson, conf)
}
