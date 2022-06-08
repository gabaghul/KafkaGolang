package main

import (
	"fmt"

	"github.com/gabaghul/golang-kafka/consts"
	"github.com/gabaghul/golang-kafka/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	log := logger.GetLogger()
	if err != nil {
		log.Panic().Msg("PANIC! couldnt load environment variables, do you have a .env file created?")
	}

	consts.GetHTTPClient()
	consts.GetTwitterAPIResources()
}

func main() {
	log := logger.GetLogger()
	resources := consts.GetTwitterAPIResources()
	log.Info().Msg(fmt.Sprintf("Bearer Token: %s", resources.BearerToken))
}
