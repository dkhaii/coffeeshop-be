package product

import (
	"errors"

	"github.com/dkhaii/cofeeshop-be/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound       = errors.New("the product was not found in the repository")
	ErrFailedToAddProduct    = errors.New("failed to add product")
	ErrFailedToUpdateProduct = errors.New("failed to update product")
	ErrFailedToDeleteProduct = errors.New("failed to delete product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(prod aggregate.Product) error
	Update(prod aggregate.Product) error
	Delete(id uuid.UUID) error
}
