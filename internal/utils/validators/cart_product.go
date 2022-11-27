package validators

import (
	"errors"

	"github.com/nuttchai/go-rest/internal/models"
)

func ValidateCartProduct(productDetail *models.Product, quantity int) error {
	var err error
	if productDetail == nil {
		err = errors.New("product not found")
	} else if productDetail.Status == "Hidden" {
		err = errors.New("product is hidden")
	} else if productDetail.Quantity < quantity {
		err = errors.New("not enough product quantity")
	}

	return err
}
