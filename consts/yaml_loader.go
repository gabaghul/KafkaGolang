package consts

import (
	"io/ioutil"
	"os"

	"github.com/gabaghul/golang-kafka/logger"
	"gopkg.in/yaml.v2"
)

type YamlLoader struct {
	TwitterV2API TwitterV2API `yaml:"twitter-v2-api"`
}

type TwitterV2API struct {
	FilteredStream FilteredStream `yaml:"filtered-stream"`
	BearerToken    string         `yaml:"bearer-token"`
}

type FilteredStream struct {
	Base  string `yaml:"base"`
	Rules string `yaml:"rules"`
}

func (yml *YamlLoader) getConfig() *YamlLoader {
	log := logger.GetLogger()
	yamlFile, err := ioutil.ReadFile("application.yml")
	if err != nil {
		log.Panic().Msgf("couldnt read application config \n %s", err.Error())
	}
	yamlFile = []byte(os.ExpandEnv(string(yamlFile))) // replaces ${var} for its values
	err = yaml.Unmarshal(yamlFile, yml)
	if err != nil {
		log.Panic().Msgf("couldnt unmarshal application config \n %s", err.Error())
	}

	log.Info().Msg("yaml successfully loaded!")

	return yml
}
