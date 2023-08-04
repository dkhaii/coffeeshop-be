package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound       = errors.New("the product was not found in the repository")
	ErrFailedToAddProduct    = errors.New("failed to add product")
	ErrFailedToUpdateProduct = errors.New("failed to update product")
	ErrFailedToDeleteProduct = errors.New("failed to delete product")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(prod Product) error
	Update(prod Product) error
	Delete(id uuid.UUID) error
}
