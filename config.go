package config

import (
	"encoding/json"
	"github.com/snormore/gologger"
	"io/ioutil"
	"path/filepath"
)

type Config interface{}

func Register(name string, c Config) error {
	return nil
}

type Configurable struct {
	Config interface{}
}

func Init(filePath string) (*Config, error) {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}
	return Read(filePath)
}

func Read(filePath string) (*Config, error) {
	logger.Info("Loading configuration from %s...", filePath)
	configJson, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return ReadJson(configJson)
}

func ReadJson(configJson []byte) (*Config, error) {
	conf := new(Config)
	err := json.Unmarshal(configJson, conf)
	return conf, err
}
