package database

import (
	"context"
	"errors"

	"github.com/fpmoles/go-microservices/internal/dberrors"
	"github.com/fpmoles/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var product []models.Product
	result := c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorId}).
		Find(&product)
	return product, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()

	result := c.DB.WithContext(ctx).
		Create(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}

		return nil, result.Error
	}

	return product, nil
}
