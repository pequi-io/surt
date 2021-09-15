package app

import (
	"github.com/surt-io/surt/pkg/antivirus"
	"github.com/surt-io/surt/pkg/antivirus/engine/clamav"
	"github.com/surt-io/surt/pkg/object"
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

	// declare empty antivirus struct
	var av antivirus.Antivirus

	// load antivirus engine
	switch cfg.Config.Antivirus.Engine {
	case "clamav":
		engine, err := clamav.New(cfg.Config.Antivirus.Network, cfg.Config.Antivirus.Address)
		if err != nil {
			log.Error().Err(err)
		}
		av = *antivirus.New(engine)

	default:
		log.Error().Msgf("antivirus %s engine is not supported", cfg.Config.Antivirus.Engine)
	}

	// test check av engine health status
	hc, err := av.GetHealthStatus()
	if err != nil {
		log.Error().Err(err)
	}

	log.Info().Msgf("health check result for %s - %s: %s", cfg.Config.Antivirus.Engine, cfg.Config.Antivirus.Address, hc)

	// test antivirus scan
	var obj object.Object
	eicarSrt := "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"
	obj.Content = []byte(eicarSrt)

	r, err := av.Scan(&obj)
	if err != nil {
		log.Error().Err(err)
	}

	log.Info().Msgf("scan result of test eicar string: %s", r)
}
