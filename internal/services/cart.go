package services

import (
	cartdto "github.com/nuttchai/go-rest/internal/dto/cart"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/validators"
)

type cartService struct {
	repo *Repository
}

type cartServiceInterface interface {
	GetCartItem(id string, cartQuery *types.CartQuery) (*models.CartItem, error)
	GetAllCartItems(userId string) ([]*models.CartItem, error)
	AddCartItem(cartDTO *cartdto.AddCartItemDTO) (*models.CartItem, error)
	UpdateCartItem(cartDTO *cartdto.UpdateCartItemDTO) (*models.CartItem, error)
	RemoveCartItem(id string) error
	PurchaseCartItem(id string) error
}

var (
	CartService cartServiceInterface
)

func init() {
	CartService = &cartService{
		repo: &repo,
	}
}

func (s *cartService) GetCartItem(id string, cartQuery *types.CartQuery) (*models.CartItem, error) {
	cartFilter := []*types.QueryFilter{}
	if cartQuery.UserId != "" {
		cartFilter = append(cartFilter, &types.QueryFilter{
			Field:    "owner_id",
			Operator: "=",
			Value:    cartQuery.UserId,
		})
	}

	return s.repo.Models.DB.GetCartItem(id, cartFilter...)
}

func (s *cartService) GetAllCartItems(userId string) ([]*models.CartItem, error) {
	return s.repo.Models.DB.GetAllCartItems(userId)
}

func (s *cartService) AddCartItem(cartDTO *cartdto.AddCartItemDTO) (*models.CartItem, error) {
	userId, productId, quantity := cartDTO.UserId, cartDTO.ProductId, cartDTO.Quantity
	productDetail, err := ProductService.GetProduct(productId)
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
	cartId, quantity := cartDTO.Id, cartDTO.Quantity
	cartDetail, err := s.repo.Models.DB.GetCartItem(cartId)
	if err != nil {
		return nil, err
	}

	productId := cartDetail.ProductId
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
	}

	productDetail, err := ProductService.GetProduct(productId)
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

func (s *cartService) PurchaseCartItem(id string) error {
	deleteResult, err := s.repo.Models.DB.RemoveCartItem(id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(deleteResult)
}
