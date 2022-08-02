package config

import (
	"log"
	"os"
	"strconv"
)

const (
	DEFAULT_DIFFICULTY = 2000000
)

func GetConfig() (string, string, int) {
	var host = DEFAULT_HOST
	if len(os.Getenv("HOST")) > 0 {
		host = os.Getenv("HOST")
	}

	var port = DEFAULT_PORT
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	var difficulty = DEFAULT_DIFFICULTY
	if len(os.Getenv("DIFFICULTY")) > 0 {
		diff, err := strconv.ParseInt(os.Getenv("DIFFICULTY"), 10, 64)
		if err != nil {
			log.Fatalf("Argument 'DIFFICULTY' should be integer, default: %d. Error: %s", DEFAULT_DIFFICULTY, err)
		}

		difficulty = int(diff)
	}

	return host, port, difficulty
}

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = "9001"
	TYPE         = "tcp"
)
