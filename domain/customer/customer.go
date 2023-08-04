package customer

import (
	"errors"
	"time"

	"github.com/dkhaii/cofeeshop-be/entity"
	"github.com/google/uuid"
)

// costum errors
var ErrInvalidPerson = errors.New("a customer must have a valid name")

type Customer struct {
	person       *entity.Person
	product      []*entity.Item
	transactions []*entity.Transaction
}

// factory to create a new costumer aggregate
func NewCustomer(name string) (Customer, error) {
	// validate if name is empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// creating new person
	person := &entity.Person{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// the new person as costumer
	return Customer{
		person:       person,
		product:      make([]*entity.Item, 0),
		transactions: make([]*entity.Transaction, 0),
	}, nil
}

// setter and getter
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) GetName() string {
	return c.person.Name
}
