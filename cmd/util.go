package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// configureLogger initialises log settings.
func configureLogger(level string) error {
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

func formatCamelCase(isArr *bool, indentCount int, s string) string {
	dash := ""
	if *isArr {
		dash = "- "
		*isArr = false
	}
	return strings.Repeat(" ", indentCount) + dash + strings.ToLower(s[:1]+s[1:])
}
