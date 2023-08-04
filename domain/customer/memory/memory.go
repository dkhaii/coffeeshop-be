package memory

import (
	"fmt"
	"sync"

	"github.com/dkhaii/cofeeshop-be/aggregate"
	customer "github.com/dkhaii/cofeeshop-be/domain/customer"
	"github.com/google/uuid"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}


// factory
func New() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mry *MemoryCustomerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, isExist := mry.customers[id]; isExist {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mry *MemoryCustomerRepository) Add(cust aggregate.Customer) error {
	if mry.customers == nil {
		mry.Lock()
		defer mry.Unlock()
		
		mry.customers = make(map[uuid.UUID]aggregate.Customer)
	}

	// check if customer is already in repo
	if _, isExist := mry.customers[cust.GetID()]; isExist {
		return fmt.Errorf("customer already exist: %w", customer.ErrFailedToAddCustomer)
	}

	mry.customers[cust.GetID()] = cust

	return nil
}

func (mry *MemoryCustomerRepository) Update(cust aggregate.Customer) error {
	// check if customer doesnt exist in repo
	if _, isExist := mry.customers[cust.GetID()]; !isExist {
		return fmt.Errorf("customer doesnt exist: %w", customer.ErrCustomerNotFound)
	}

	mry.Lock()
	defer mry.Unlock()
	
	mry.customers[cust.GetID()] = cust

	return nil
}
