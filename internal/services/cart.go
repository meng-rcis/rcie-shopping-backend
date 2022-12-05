package services

import (
	"errors"

	cartdto "github.com/nuttchai/go-rest/internal/dto/cart"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type cartService struct {
	repo *Repository
}

type cartServiceInterface interface {
	GetAllCartItems(userId string) ([]*models.CartItem, error)
	AddCartItem(cartDTO *cartdto.AddCartItemDTO) (*models.CartItem, error)
	UpdateCartItem(cartDTO *cartdto.UpdateCartItemDTO) (*models.CartItem, error)
	RemoveCartItem(id string) error
}

var (
	CartService cartServiceInterface
)

func init() {
	CartService = &cartService{
		repo: &repo,
	}
}

func (s *cartService) GetAllCartItems(userId string) ([]*models.CartItem, error) {
	return s.repo.Models.DB.GetAllCartItems(userId)
}

func (s *cartService) AddCartItem(cartDTO *cartdto.AddCartItemDTO) (*models.CartItem, error) {
	userId := cartDTO.UserId
	productId := cartDTO.ProductId
	quantity := cartDTO.Quantity
	productDetail, err := ProductService.GetProductDetail(productId)
	if err != nil {
		return nil, err
	} else if err = validators.ValidateCartItem(
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

	return s.repo.Models.DB.AddCartItem(
		userId,
		productId,
		quantity,
		float64(quantity)*productDetail.Price,
	)
}

func (s *cartService) UpdateCartItem(cartDTO *cartdto.UpdateCartItemDTO) (*models.CartItem, error) {
	cartId := cartDTO.Id
	productId := cartDTO.ProductId
	quantity := cartDTO.Quantity
	if quantity < 1 {
		return nil, errors.New("updated quantity must be greater than 0")
	}

	cartDetail, err := s.repo.Models.DB.GetCartItem(cartId)
	if err != nil {
		return nil, err
	}

	if cartDetail.Quantity < quantity {
		quantityDiff := quantity - cartDetail.Quantity
		if err = ProductService.DeductProductQuantity(
			productId,
			quantityDiff,
		); err != nil {
			return nil, err
		}
	} else if cartDetail.Quantity > quantity {
		quantityDiff := cartDetail.Quantity - quantity
		if err = ProductService.AddProductQuantity(
			productId,
			quantityDiff,
		); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("cannot update cart item with same quantity")
	}

	productDetail, err := ProductService.GetProductDetail(productId)
	if err != nil {
		return nil, err
	}

	return s.repo.Models.DB.UpdateCartItem(
		cartId,
		quantity,
		float64(quantity)*productDetail.Price,
	)
}

func (s *cartService) RemoveCartItem(id string) error {
	cartDetail, err := s.repo.Models.DB.GetCartItem(id)
	if err != nil {
		return err
	}

	deleteResult, err := s.repo.Models.DB.RemoveCartItem(id)
	if err != nil {
		return err
	}

	if err = ProductService.AddProductQuantity(
		cartDetail.ProductId,
		cartDetail.Quantity,
	); err != nil {
		return err
	}

	return validators.CheckRowsAffected(deleteResult)
}
