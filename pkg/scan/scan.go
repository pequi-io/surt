package scan

import (
	"fmt"
	"time"

	"github.com/surt-io/surt/internal/config"
	"github.com/surt-io/surt/internal/util"
	"github.com/surt-io/surt/pkg/antivirus"
	"github.com/surt-io/surt/pkg/antivirus/engine/clamav"
	"github.com/surt-io/surt/pkg/types"
)

type Service struct {
	repo Repository
	av   antivirus.Antivirus
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) SetupEngine(cfg config.Antivirus) (err error) {

	switch cfg.Engine {
	case "clamav":
		engine, err := clamav.New(cfg.Network, cfg.Address)
		if err != nil {
			return err
		}
		s.av = *antivirus.New(engine)
		return nil
	default:
		return fmt.Errorf("antivirus %s engine is not supported", cfg.Engine)
	}

}

func (s *Service) CreateScan(rawFilePath string) (types.ID, error) {
	n, err := types.NewScan(rawFilePath)
	if err != nil {
		return n.ID, err
	}
	return s.repo.Create(n)
}

func (s *Service) UpdateScan(sc *types.Scan) (err error) {
	return s.repo.Update(sc)
}

func (s *Service) GetScan(id types.ID) (*types.Scan, error) {
	return s.repo.Get(id)
}

func (s *Service) ExecuteScan(sc *types.Scan) (err error) {

	// Update Scan status
	sc.StartedAt = time.Now()
	sc.Status = "RUNNING"
	err = s.UpdateScan(sc)
	if err != nil {
		return err
	}

	// Execute Scan
	r, err := s.av.ScanObject(&sc.Object)
	if err != nil {
		return err
	}

	// Add result to Scan
	sc.Result = r
	sc.FinishedAt = time.Now()

	// Parse status
	status := util.ParseScanStatus(r[0].Status)
	if status == "INFECTED" {
		sc.Infected = true
	} else {
		sc.Infected = false
	}

	err = s.UpdateScan(sc)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) HealthCheck() (string, error) {
	return s.av.GetHealthStatus()
}
