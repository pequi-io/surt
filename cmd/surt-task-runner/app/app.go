package app

import (
	"fmt"
	"sync"

	"github.com/surt-io/surt/internal/config"
	"github.com/surt-io/surt/internal/healthz"
	"github.com/surt-io/surt/internal/logger"
	"github.com/surt-io/surt/pkg/antivirus"
	"github.com/surt-io/surt/pkg/antivirus/engine/clamav"
	"github.com/surt-io/surt/pkg/types"
)

// define log with new logger
var log = logger.New()

// create cfg config.File type
var cfg *config.File

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

func RunApp() {

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		RunTaskRunner()
		wg.Done()
	}()

	go func(p string) {
		RunHealthz(p)
		wg.Done()
	}(cfg.Config.API.Port)

	wg.Wait()
}

func RunHealthz(port string) (err error) {
	h := healthz.New()
	log.Info().Msg("starting healthcheck...")
	err = h.Run(":" + port)

	if err != nil {
		return
	}

	return
}

func RunTaskRunner() (err error) {

	// declare empty antivirus struct
	var av antivirus.Antivirus

	// load antivirus engine
	switch cfg.Config.Antivirus.Engine {
	case "clamav":
		engine, err := clamav.New(cfg.Config.Antivirus.Network, cfg.Config.Antivirus.Address)
		if err != nil {
			log.Error().Err(err)
			return err
		}
		av = *antivirus.New(engine)

	default:
		err = fmt.Errorf("antivirus %s engine is not supported", cfg.Config.Antivirus.Engine)
		log.Error().Err(err)
		return
	}

	// test check av engine health status
	hc, err := av.GetHealthStatus()
	if err != nil {
		log.Error().Err(err)
		return
	}

	log.Info().Msgf("health check result for %s - %s: %s", cfg.Config.Antivirus.Engine, cfg.Config.Antivirus.Address, hc)

	// test antivirus scan
	var obj types.Object
	eicarSrt := "clean"
	//eicarSrt := "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"
	obj.Content = []byte(eicarSrt)

	r, err := av.Scan(&obj)
	if err != nil {
		log.Error().Err(err)
		return
	}

	log.Info().Msgf("scan result of test eicar string: %v", r)

	return nil
}
