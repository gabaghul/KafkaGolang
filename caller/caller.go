package caller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gabaghul/golang-kafka/consts"
	"github.com/gabaghul/golang-kafka/logger"
	"github.com/gabaghul/golang-kafka/model"
)

func SetRules() {
	log := logger.GetLogger()
	client := consts.GetHTTPClient()

	log.Info().Msg("starting set rules process")

	apiResources := consts.GetTwitterAPIResources()
	contentType := "application/json"

	ruleRequest := model.SetStreamRuleRequest{
		Add: consts.GetSetRules(),
	}

	payload, err := json.Marshal(&ruleRequest)
	if err != nil {
		log.Err(err).Msg("couldnt marshal set rules payload")
		return
	}

	req, err := http.NewRequest("POST", apiResources.FilteredStream.Rules, bytes.NewBuffer(payload))
	if err != nil {
		log.Err(err).Msg("couldnt create request for twitter set stream rules api")
		return
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiResources.BearerToken))

	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("couldnt post into twitter set stream rules api")
		return
	}
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		log.Info().Msg("rules successfully updated")
		return
	}

	log.Warn().Msg(fmt.Sprint("rules didnt update as expected, see returned status code: ", resp.StatusCode))
}

func GetRules() (res model.GetStreamRuleResponse) {
	log := logger.GetLogger()
	client := consts.GetHTTPClient()

	log.Info().Msg("starting get rules process")

	apiResources := consts.GetTwitterAPIResources()

	req, err := http.NewRequest("GET", apiResources.FilteredStream.Rules, nil)
	if err != nil {
		log.Err(err).Msg("couldnt create request for twitter get stream rules api")
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiResources.BearerToken))

	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("couldnt post into twitter get stream rules api")
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Warn().Msg(fmt.Sprint("didnt get rules as expected, see returned status code: ", resp.StatusCode))
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		log.Err(err).Msg("couldnt marshall get stream rule json response")
	}

	return
}

func DeleteAllRules(IDs []string) {
	if len(IDs) == 0 {
		return
	}

	log := logger.GetLogger()
	client := consts.GetHTTPClient()

	log.Info().Msg("starting delete rules process")

	apiResources := consts.GetTwitterAPIResources()
	contentType := "application/json"

	ruleRequest := model.DeleteStreamRuleRequest{
		Delete: model.StremRuleDelete{
			IDs: IDs,
		},
	}

	payload, err := json.Marshal(&ruleRequest)
	if err != nil {
		log.Err(err).Msg("couldnt marshal delete rule payload")
		return
	}

	req, err := http.NewRequest("POST", apiResources.FilteredStream.Rules, bytes.NewBuffer(payload))
	if err != nil {
		log.Err(err).Msg("couldnt create request for twitter set stream rules api")
		return
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiResources.BearerToken))

	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("couldnt post into twitter delete stream rules api")
		return
	}
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		log.Info().Msg("rules successfully deleted")
		return
	}

	log.Warn().Msg(fmt.Sprint("rules didnt delete as expected, see returned status code: ", resp.StatusCode))
}

func GetTwitterStream() {
	log := logger.GetLogger()
	client := consts.GetHTTPClient()

	log.Info().Msg("starting get streaming process")

	apiResources := consts.GetTwitterAPIResources()

	req, err := http.NewRequest("GET", apiResources.FilteredStream.Base, nil)
	if err != nil {
		log.Err(err).Msg("couldnt create request for twitter get stream api")
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiResources.BearerToken))

	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("couldnt post into twitter get stream api")
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Error().Msg(fmt.Sprint("didnt get stream as expected, see returned status code: ", resp.StatusCode))
		return
	}
	reader := bufio.NewReader(resp.Body)
	i := 1
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			log.Info().Msg("stream finished, terminating process successfully")
			return
		}
		if err != nil {
			if i > 5 {
				log.Err(err).Msg("retries exceeded, terminating process")
				return
			}
			log.Err(err).Msg(fmt.Sprintf("something went wrong on stream buffer, attempting for the %d time", i))
			i++
			continue
		}
		payload := string(line)
		if len(payload) > 2 {
			log.Info().Msg(payload)
		}
		// res :=
		// err = json.Unmarshal(line, &res)
		// if err != nil {
		// 	log.Err(err).Msg("couldnt marshall get stream rule json response")
		// }
	}
}
