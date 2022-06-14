package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gabaghul/golang-kafka/caller"
	"github.com/gabaghul/golang-kafka/consts"
	"github.com/gabaghul/golang-kafka/logger"
	"github.com/joho/godotenv"
)

func init() {
	log := logger.GetLogger()
	err := godotenv.Load()
	if err != nil {
		log.Panic().Msg("PANIC! couldnt load environment variables, do you have a .env file created?")
	}

	consts.GetHTTPClient()
	consts.GetTwitterAPIResources()
}

func main() {
	log := logger.GetLogger()

	env := os.Getenv("ENV")
	setRules, err := strconv.ParseBool(os.Getenv("SET_STREAM_RULES"))
	if err != nil {
		log.Err(err).Msg("couldnt define if it'll set stream rules")
		return
	}
	rules := caller.GetRules()
	if setRules {
		ids := make([]string, len(rules.Data))
		for i, rule := range rules.Data {
			ids[i] = rule.ID
		}
		caller.DeleteAllRules(ids)
		caller.SetRules()
		rules = caller.GetRules()
	}
	if env == "dev" {
		log.Debug().Msg(fmt.Sprint("RESPONSE: ", rules))
	}

	caller.GetTwitterStream()
}
