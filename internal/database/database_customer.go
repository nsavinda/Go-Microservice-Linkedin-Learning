package database

import (
	"context"
	"errors"

	"github.com/fpmoles/go-microservices/internal/dberrors"
	"github.com/fpmoles/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customer []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAddress}).
		Find(&customer)
	return customer, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()

	result := c.DB.WithContext(ctx).
		Create(&customer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}

		return nil, result.Error

	}

	return customer, nil

}