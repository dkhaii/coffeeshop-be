package aggregate

import (
	"errors"

	"github.com/dkhaii/cofeeshop-be/entity"
	"github.com/dkhaii/cofeeshop-be/valueobject"
	"github.com/google/uuid"
)

var ErrInvalidPerson = errors.New("a costumer must have a valid name")

type Costumer struct {
	person       *entity.Person
	product      []*entity.Item
	transactions []*valueobject.Transaction
}

// factory to create a new costumer aggregate
func NewCustomer(name string) (Costumer, error) {
	if name == "" {
		return Costumer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Costumer{
		person:       person,
		product:      make([]*entity.Item, 0),
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}
