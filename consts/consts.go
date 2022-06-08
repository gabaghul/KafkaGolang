package consts

import (
	"net/http"
	"sync"
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
