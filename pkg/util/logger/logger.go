package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func New(isDebug bool, isPretty bool) zerolog.Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	logger := zerolog.New(os.Stderr).Level(logLevel).With().Timestamp().Logger()
	if isPretty {
		logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	return logger
}

func NewDefault() zerolog.Logger {
	logLevel := zerolog.DebugLevel
	logger := zerolog.New(os.Stderr).Level(logLevel).With().Timestamp().Logger()
	return logger
}
