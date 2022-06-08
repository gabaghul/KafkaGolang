package caller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gabaghul/golang-kafka/consts"
	"github.com/gabaghul/golang-kafka/logger"
	"github.com/gabaghul/golang-kafka/model"
)

func GetRules() {
	log := logger.GetLogger()
	client := consts.GetHTTPClient()

	apiResources := consts.GetTwitterAPIResources()
	contentType := "application/json"

	ruleRequest := model.StreamModelRequest{
		Add: []model.StreamModel{
			{
				Value: "dog has:images",
				Tag:   "dog pictures",
			},
		},
	}

	payload, err := json.Marshal(&ruleRequest)
	if err != nil {
		log.Err(err).Msg("couldnt marshal payload")
		return
	}

	req, err := http.NewRequest("POST", apiResources.FilteredStream.Rules, bytes.NewBuffer(payload))
	if err != nil {
		log.Err(err).Msg("couldnt create request for twitter stream rules api")
		return
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiResources.BearerToken))

	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("couldnt post into twitter stream rules api")
		return
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	fmt.Println(string(bodyBytes), resp.StatusCode)
}
