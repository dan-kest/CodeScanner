package services

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/interfaces"
	"github.com/dan-kest/cscanner/internal/models"
)

type ScanService struct {
	conf           *config.Config
	scanRepository interfaces.ScanRepository
}

func NewScanService(conf *config.Config, scanRepository interfaces.ScanRepository) *ScanService {
	return &ScanService{
		conf:           conf,
		scanRepository: scanRepository,
	}
}

func (s *ScanService) RunTask(task *models.Task) {

}

func (s *ScanService) RunErrorTask(body []byte, err error) error {
	return s.scanRepository.CreateTaskErrorLog(body, err)
}
