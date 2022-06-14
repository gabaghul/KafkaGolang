package consts

import (
	"net/http"
	"sync"

	"github.com/gabaghul/golang-kafka/model"
)

var (
	client     *http.Client
	yamlLoader *YamlLoader
	once       sync.Once
)

func GetHTTPClient() *http.Client {
	once.Do(func() {
		client = &http.Client{}
	})

	return client
}

func GetApplicationResources() Application {
	if yamlLoader == nil {
		yamlLoader = &YamlLoader{}
		yamlLoader.getConfig()
	}

	return yamlLoader.Application
}

func GetTwitterAPIResources() TwitterV2API {
	if yamlLoader == nil {
		yamlLoader = &YamlLoader{}
		yamlLoader.getConfig()
	}

	return yamlLoader.TwitterV2API
}

func GetSetRules() []model.SetStreamRule {
	return []model.SetStreamRule{
		{
			Value: "dog has:images",
			Tag:   "dog pictures",
		},
		{
			Value: "cat has:images",
			Tag:   "cat pictures",
		},
	}
}
