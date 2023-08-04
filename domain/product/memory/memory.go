package memory

import (
	"fmt"
	"sync"

	"github.com/dkhaii/cofeeshop-be/aggregate"
	"github.com/dkhaii/cofeeshop-be/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mry *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	// convert the map into a slice
	for _, product := range mry.products {
		products = append(products, product)
	}

	return products, nil
}

func (mry *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, isExist := mry.products[id]; isExist {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (mry *MemoryProductRepository) Add(prod aggregate.Product) error {
	mry.Lock()
	defer mry.Unlock()

	if _, isExist := mry.products[prod.GetID()]; isExist {
		return fmt.Errorf("product already exist: %w", product.ErrFailedToAddProduct)
	}

	mry.products[prod.GetID()] = prod

	return nil
}

func (mry *MemoryProductRepository) Update(prod aggregate.Product) error {
	mry.Lock()
	defer mry.Unlock()

	if _, isExist := mry.products[prod.GetID()]; !isExist {
		return fmt.Errorf("product doesnt exist: %w", product.ErrProductNotFound)
	}

	mry.products[prod.GetID()] = prod

	return nil
}

func (mry *MemoryProductRepository) Delete(id uuid.UUID) error {
	mry.Lock()
	defer mry.Unlock()

	if _, isExist := mry.products[id]; !isExist {
		return fmt.Errorf("prodcut doesnt exist: %w", product.ErrProductNotFound)
	}

	delete(mry.products, id)

	return nil
}
