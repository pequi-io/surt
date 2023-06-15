package logger

import (
	"os"

	"github.com/pequi-io/surt/pkg/config"
	"github.com/rs/zerolog"
)

// create cfg config.File type
var cfg *config.File

// default log level
var logLevel = zerolog.DebugLevel

// default logger
var log = zerolog.New(os.Stderr).Level(logLevel).With().Timestamp().Logger()

// init function loads log config
func init() {

	var err error

	// create cfg with default values
	cfg, err = config.Default()
	if err != nil {
		log.Error().Err(err)
	}

	// load values from config file
	err = cfg.LoadConfig()
	if err != nil {
		log.Error().Err(err)
	}
}

func New() *zerolog.Logger {
	logLevel := zerolog.InfoLevel
	if cfg.Config.Log.Debug {
		logLevel = zerolog.DebugLevel
	}
	logger := zerolog.New(os.Stderr).Level(logLevel).With().Timestamp().Logger()
	if cfg.Config.Log.Pretty {
		log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	return &logger
}

func NewDefault() *zerolog.Logger {
	return &log
}
