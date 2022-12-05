package validators

import (
	"errors"

	"github.com/nuttchai/go-rest/internal/models"
)

func ValidateCartItem(productDetail *models.Product, quantity int) error {
	if productDetail == nil {
		return errors.New("product not found")
	} else if productDetail.Status == "Hidden" {
		return errors.New("product is hidden")
	} else if productDetail.Quantity < quantity {
		return errors.New("product quantity is not enough")
	}
	return nil
}
