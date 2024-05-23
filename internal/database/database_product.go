package database

import (
	"context"

	"github.com/fpmoles/go-microservices/internal/models"
)

func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var product []models.Product
	result := c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorId}).
		Find(&product)
	return product, result.Error
}
