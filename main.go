package main

import (
	"github.com/gabaghul/golang-kafka/caller"
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
	caller.GetRules()
}
