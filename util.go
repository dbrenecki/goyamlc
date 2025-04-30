package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ConfigureLogger initialises log settings.
func ConfigureLogger(level string) error {
	zerolog.TimeFieldFormat = time.RFC822
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		return fmt.Errorf(
			`error: "--log-level %s is not a valid level. Valid log levels ["info", "debug", "warn", "error", "fatal"]`,
			logLevel.String())

	}
	zerolog.SetGlobalLevel(logLevel)
	log.Info().Msgf("setting log-level: %#v", logLevel.String())
	return nil
}
