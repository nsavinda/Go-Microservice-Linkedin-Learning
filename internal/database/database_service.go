package database

import (
	"context"

	"github.com/fpmoles/go-microservices/internal/models"
)

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var service []models.Service
	result := c.DB.WithContext(ctx).
		Find(&service)

	return service, result.Error
}
