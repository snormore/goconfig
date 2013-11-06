package config

import (
	"encoding/json"
	"github.com/snormore/gologger"
	"io/ioutil"
	"path/filepath"
	"reflect"
)

type Config interface{}

func Register(name string, c Config) error {
	return nil
}

type Configurable struct {
	Config Config
}

func Init(conf *Config, filePath string) error {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	return Read(conf, filePath)
}

func Read(conf *Config, filePath string) error {
	logger.Info("Loading configuration from %s...", filePath)
	configJson, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return ReadJson(conf, configJson)
}

func ReadJson(conf *Config, configJson []byte) error {
	typedConfig := *conf
	logger.Info("conf = %+v, %+v", typedConfig, reflect.TypeOf(typedConfig))
	err := json.Unmarshal(configJson, &typedConfig)
	logger.Info("conf = %+v, %+v", typedConfig, reflect.TypeOf(typedConfig))
	// logger.Info("verbosity = %+v", typedConfig["Verbosity"])
	conf = &typedConfig
	return err
}
