package app

import (
	"sync"

	"github.com/surt-io/surt/internal/config"
	"github.com/surt-io/surt/internal/healthz"
	"github.com/surt-io/surt/internal/logger"
	"github.com/surt-io/surt/pkg/repository"
	"github.com/surt-io/surt/pkg/scan"
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

	repo := repository.NewScanDynamoDB("surt_scan")
	svc := scan.NewService(repo)

	// setup av engine
	err = svc.SetupEngine(cfg.Config.Antivirus)
	if err != nil {
		log.Error().Err(err)
		return
	}

	// test check av engine health status
	hc, err := svc.HealthCheck()
	if err != nil {
		log.Error().Err(err)
		return
	}

	log.Info().Msgf("health check result for %s - %s: %s", cfg.Config.Antivirus.Engine, cfg.Config.Antivirus.Address, hc)

	//create new scan
	log.Info().Msg("creating new scan task")
	scanId, err := svc.CreateScan("s3://mybucket/eicar.zip")
	if err != nil {
		log.Err(err)
		return
	}

	log.Info().Msgf("New ScanID: %v", scanId)

	s, err := svc.GetScan(scanId)
	if err != nil {
		log.Err(err)
		return
	}

	log.Info().Msgf("Get Scan: %v", s)

	log.Info().Msg("Execute scan")
	err = svc.ExecuteScan(s)
	if err != nil {
		log.Err(err)
		return
	}

	log.Info().Msgf("Object infected: %v", s.Infected)

	return nil
}
