package services

import (
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/types"
)

type orderService struct {
	repo *Repository
}

type orderServiceInterface interface {
	GetOrders(orderQuery *types.OrderQuery) ([]*models.Order, error)
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
	orderFilter := []*types.QueryFilter{}
	if orderQuery.UserId != "" {
		orderFilter = append(orderFilter, &types.QueryFilter{
			Field: "o.owner_id",
			Value: orderQuery.UserId,
		})
	}

	if orderQuery.Status != "" {
		orderFilter = append(orderFilter, &types.QueryFilter{
			Field: "os.name",
			Value: orderQuery.Status,
		})
	}

	return s.repo.Models.DB.GetOrders(orderFilter...)
}
