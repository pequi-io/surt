package scan

import (
	"fmt"
	"time"

	"github.com/pequi-io/surt/pkg/antivirus"
	"github.com/pequi-io/surt/pkg/antivirus/engine/clamav"
	"github.com/pequi-io/surt/pkg/config"
	"github.com/pequi-io/surt/pkg/entity"
	"github.com/pequi-io/surt/pkg/util"
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

func (s *Service) CreateScan(rawFilePath string) (entity.ID, error) {
	n, err := entity.NewScan(rawFilePath)
	if err != nil {
		return n.ID, err
	}
	return s.repo.Create(n)
}

func (s *Service) UpdateScan(sc *entity.Scan) (err error) {
	return s.repo.Update(sc)
}

func (s *Service) GetScan(id entity.ID) (*entity.Scan, error) {
	return s.repo.Get(id)
}

func (s *Service) ExecuteScan(sc *entity.Scan) (err error) {

	// Update Scan status
	sc.StartedAt = time.Now()
	sc.Status = "RUNNING"
	err = s.UpdateScan(sc)
	if err != nil {
		return err
	}

	// TO-DO
	// retrieve object from storage service (GetObject)

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

	// TO-DO
	// update object tags with scan result

	return nil

}

func (s *Service) HealthCheck() (string, error) {
	return s.av.GetHealthStatus()
}
