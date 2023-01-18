package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type productService struct {
	repo *Repository
}

type productServiceInterface interface {
	GetProduct(id string) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	AddProductQuantity(id string, quantity int) error
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

func (s *productService) GetProduct(id string) (*models.Product, error) {
	return s.repo.Models.DB.GetProduct(id)
}

func (s *productService) UpdateProduct(product *models.Product) (*models.Product, error) {
	updatedProduct, err := s.repo.Models.DB.UpdateProduct(product)

	return updatedProduct, err
}

func (s *productService) AddProductQuantity(id string, quantity int) error {
	result, err := s.repo.Models.DB.AddProductQuantity(id, quantity)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}

func (s *productService) DeductProductQuantity(id string, quantity int) error {
	result, err := s.repo.Models.DB.DeductProductQuantity(id, quantity)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
