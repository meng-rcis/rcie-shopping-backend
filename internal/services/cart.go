package services

import "github.com/nuttchai/go-rest/internal/models"

type cartService struct {
	repo *Repository
}

type cartServiceInterface interface {
	GetItems(userId string) ([]*models.CartItem, error)
}

var (
	CartService cartServiceInterface
)

func init() {
	CartService = &cartService{
		repo: &repo,
	}
}

func (s *cartService) GetItems(userId string) ([]*models.CartItem, error) {
	return s.repo.Models.DB.GetItems(userId)
}
