package main

import (
	"log"

	"github.com/michaelpeterswa/go-start/internal/logging"
)

func main() {
	logger, err := logging.InitZap()
	if err != nil {
		log.Panicf("could not acquire zap logger: %s", err.Error())
	}
	logger.Info("go-start init...")
}
