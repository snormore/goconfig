package config

import (
	"github.com/snormore/gologger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	logger.Verbosity = 2
}

type SampleConfig struct {
	AppName   string `json:"app_name"`
	Verbosity int    `json:"verbosity"`
}

const SampleConfigJson = `
{
  "app_name": "sample-app",
  "verbosity": 1
}
`

var SampleConfigObject = SampleConfig{
	AppName:   "sample-app",
	Verbosity: 1,
}

const SampleConfigJson2 = `
{
  "app_name": "sample-app-2",
  "verbosity": 2
}
`

var SampleConfigObject2 = SampleConfig{
	AppName:   "sample-app-2",
	Verbosity: 2,
}

func TestReadJson(t *testing.T) {
	conf := Config(new(SampleConfig))
	err := ReadJson([]byte(SampleConfigJson), &conf)
	assert.NoError(t, err)
	assert.Equal(t, SampleConfigObject.AppName, conf.(*SampleConfig).AppName)
	assert.Equal(t, SampleConfigObject.Verbosity, conf.(*SampleConfig).Verbosity)
	assert.NotEqual(t, SampleConfigObject2.AppName, conf.(*SampleConfig).AppName)
	assert.NotEqual(t, SampleConfigObject2.Verbosity, conf.(*SampleConfig).Verbosity)
}
