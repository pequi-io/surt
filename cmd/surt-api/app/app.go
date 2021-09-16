package app

import (
	"github.com/surt-io/surt/internal/config"
	"github.com/surt-io/surt/internal/logger"
	"github.com/surt-io/surt/pkg/api"
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
	r := api.New()
	r.Run(":" + cfg.Config.API.Port)
}
