package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type cartService struct {
	repo *Repository
}

type cartServiceInterface interface {
	GetAllCartProducts(userId string) ([]*models.CartItem, error)
	AddCartProduct(userId string, productId string, quantity int) (*models.CartItem, error)
}

var (
	CartService cartServiceInterface
)

func init() {
	CartService = &cartService{
		repo: &repo,
	}
}

func (s *cartService) GetAllCartProducts(userId string) ([]*models.CartItem, error) {
	return s.repo.Models.DB.GetAllCartProducts(userId)
}

func (s *cartService) AddCartProduct(userId string, productId string, quantity int) (*models.CartItem, error) {
	productDetail, err := ProductService.GetProductDetail(productId)
	if err != nil {
		return nil, err
	} else if err = validators.ValidateCartProduct(
		productDetail,
		quantity,
	); err != nil {
		return nil, err
	} else if err = ProductService.DeductProductQuantity(
		productId,
		quantity,
	); err != nil {
		return nil, err
	}

	return s.repo.Models.DB.AddCartProduct(
		userId,
		productId,
		quantity,
		float64(quantity)*productDetail.Price,
	)
}
