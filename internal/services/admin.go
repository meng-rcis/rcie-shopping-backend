package services

import (
	admindto "github.com/nuttchai/go-rest/internal/dto/admin"
	"github.com/nuttchai/go-rest/internal/models"
)

type adminService struct {
	repo *Repository
}

type adminServiceInterface interface {
	UpdateOrderStatus(updatedOrderDto *admindto.UpdateOrderStatusDTO) (*models.Order, error)
	AddProductQuantity(updatedProductDto *admindto.AddProductQuantityDTO) error
}

var (
	AdminService adminServiceInterface
)

func init() {
	AdminService = &adminService{
		repo: &repo,
	}
}

func (s *adminService) UpdateOrderStatus(updatedOrderDto *admindto.UpdateOrderStatusDTO) (*models.Order, error) {
	return OrderService.UpdateOrder(
		&models.Order{
			Id:     updatedOrderDto.OrderId,
			Status: updatedOrderDto.Status,
		},
	)
}

func (s *adminService) AddProductQuantity(updatedProductDto *admindto.AddProductQuantityDTO) error {
	return ProductService.AddProductQuantity(
		updatedProductDto.ProductId,
		updatedProductDto.Quantity,
	)
}
