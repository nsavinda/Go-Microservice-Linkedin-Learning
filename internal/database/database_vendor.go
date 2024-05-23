package database

import (
	"context"

	"github.com/fpmoles/go-microservices/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendor []models.Vendor
	result := c.DB.WithContext(ctx).
		Find(&vendor)

	return vendor, result.Error
}
