package app

import (
	"github.com/pequi-io/surt/pkg/api"
	"github.com/pequi-io/surt/pkg/config"
	"github.com/pequi-io/surt/pkg/logger"
)

// define log with new logger
var log = logger.New()

// create cfg config.File type
var cfg *config.File

func init() {

	var err error

	// populate cfg with default values
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

func RunApp() {
	log.Info().Msg("starting api server on port " + cfg.Config.API.Port)
	s := api.New()
	err := s.Run(":" + cfg.Config.API.Port)
	if err != nil {
		log.Fatal().Err(err)
	}
}
