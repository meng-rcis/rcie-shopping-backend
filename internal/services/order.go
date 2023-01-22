package services

import (
	orderdto "github.com/nuttchai/go-rest/internal/dto/order"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/query"
)

type orderService struct {
	repo *Repository
}

type orderServiceInterface interface {
	GetOrders(orderQuery *types.OrderQuery) ([]*models.Order, error)
	CreateOrder(order *orderdto.CreateOrderDTO) (*models.Order, error)
	UpdateOrder(order *models.Order) (*models.Order, error)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{
		repo: &repo,
	}
}

func (s *orderService) GetOrders(orderQuery *types.OrderQuery) ([]*models.Order, error) {
	filters := query.GenerateOrderFilter(orderQuery)
	return s.repo.Models.DB.GetOrders(filters...)
}

func (s *orderService) CreateOrder(orderdto *orderdto.CreateOrderDTO) (*models.Order, error) {
	cartId, userId := orderdto.CartId, orderdto.UserId
	cart, err := CartService.GetCartItem(
		cartId,
		&types.CartQuery{
			UserId: userId,
		},
	)
	if err != nil {
		return nil, err
	}

	newOrder, err := s.repo.Models.DB.CreateOrder(
		&models.Order{
			OwnerId:    cart.OwnerId,
			ProductId:  cart.ProductId,
			Quantity:   cart.Quantity,
			TotalPrice: cart.TotalPrice,
		},
	)
	if err != nil {
		return nil, err
	}

	err = CartService.PurchaseCartItem(cartId)
	if err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (s *orderService) UpdateOrder(orderdto *models.Order) (*models.Order, error) {
	// TODO: Check if user status is valid to edit order status
	return s.repo.Models.DB.UpdateOrder(orderdto)
}
