package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type sampleService struct {
	repo *Repository
}

type sampleServiceInterface interface {
	GetSample(id string) (*models.Sample, error)
	CreateSample(sample *models.NewSample) (int, error)
	UpdateSample(sample *models.Sample) error
	DeleteSample(id string) error
	EmptySampleDesc(id string) error
}

var (
	SampleService sampleServiceInterface
)

func init() {
	SampleService = &sampleService{
		repo: &repo,
	}
}

func (s *sampleService) GetSample(id string) (*models.Sample, error) {
	return s.repo.Models.DB.GetSample(id)
}

func (s *sampleService) CreateSample(sample *models.NewSample) (int, error) {
	return s.repo.Models.DB.CreateSample(sample)
}

func (s *sampleService) UpdateSample(sample *models.Sample) error {
	result, err := s.repo.Models.DB.UpdateSample(sample)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}

func (s *sampleService) DeleteSample(id string) error {
	result, err := s.repo.Models.DB.DeleteSample(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}

func (s *sampleService) EmptySampleDesc(id string) error {
	result, err := s.repo.Models.DB.EmptySampleDesc(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
