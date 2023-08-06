package memory

import (
	"fmt"
	"sync"

	"github.com/dkhaii/cofeeshop-be/domain/product"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mry *MemoryRepository) GetAll() ([]product.Product, error) {
	var products []product.Product

	// convert the map into a slice
	for _, product := range mry.products {
		products = append(products, product)
	}

	return products, nil
}

func (mry *MemoryRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, isExist := mry.products[id]; isExist {
		return product, nil
	}

	return product.Product{}, product.ErrProductNotFound
}

func (mry *MemoryRepository) Add(prod product.Product) error {
	mry.Lock()
	defer mry.Unlock()

	if _, isExist := mry.products[prod.GetID()]; isExist {
		return fmt.Errorf("product already exist: %w", product.ErrFailedToAddProduct)
	}

	mry.products[prod.GetID()] = prod

	return nil
}

func (mry *MemoryRepository) Update(prod product.Product) error {
	mry.Lock()
	defer mry.Unlock()

	if _, isExist := mry.products[prod.GetID()]; !isExist {
		return fmt.Errorf("product doesnt exist: %w", product.ErrProductNotFound)
	}

	mry.products[prod.GetID()] = prod

	return nil
}

func (mry *MemoryRepository) Delete(id uuid.UUID) error {
	mry.Lock()
	defer mry.Unlock()

	if _, isExist := mry.products[id]; !isExist {
		return fmt.Errorf("prodcut doesnt exist: %w", product.ErrProductNotFound)
	}

	delete(mry.products, id)

	return nil
}
