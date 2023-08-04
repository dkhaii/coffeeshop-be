package aggregate

import (
	"errors"
	"time"

	"github.com/dkhaii/cofeeshop-be/entity"
	"github.com/google/uuid"
)

var ErrMissingValues = errors.New("missing an important values")

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name string, price float64, qty int) (Product, error) {
	if name == "" || price == 0 || qty == 0 {
		return Product{}, ErrMissingValues
	}

	product := &entity.Item{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return Product{
		item:     product,
		price:    price,
		quantity: qty,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
