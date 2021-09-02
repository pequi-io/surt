package app

import (
	"github.com/surt-io/surt/pkg/apis"
	"github.com/surt-io/surt/pkg/util/config"
	"github.com/surt-io/surt/pkg/util/logger"
)

func RunApp() {

	// define log with new logger default
	log := logger.NewDefault()

	// create cfg with default values
	cfg, err := config.Default()
	if err != nil {
		log.Error().Err(err)
	}

	// load values from config file
	err = cfg.LoadConfig()
	if err != nil {
		log.Error().Err(err)
	}

	// update logger with global log config values
	log = logger.New(cfg.Config.Log.Debug, cfg.Config.Log.Debug)

	log.Info().Msg("starting api server on port " + cfg.Config.API.Port)
	r := apis.SetupRouter()
	r.Run(":" + cfg.Config.API.Port)
}
