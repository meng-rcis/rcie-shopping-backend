package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type productService struct {
	repo *Repository
}

type productServiceInterface interface {
	GetProductDetail(id string) (*models.Product, error)
	DeductProductQuantity(id string, quantity int) error
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{
		repo: &repo,
	}
}

func (s *productService) GetProductDetail(id string) (*models.Product, error) {
	return s.repo.Models.DB.GetProductDetail(id)
}

func (s *productService) DeductProductQuantity(id string, quantity int) error {
	result, err := s.repo.Models.DB.DeductProductQuantity(id, quantity)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
