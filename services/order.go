package services

import (
	"log"

	"github.com/dkhaii/cofeeshop-be/aggregate"
	costumer "github.com/dkhaii/cofeeshop-be/domain/customer"
	"github.com/dkhaii/cofeeshop-be/domain/customer/memory"
	"github.com/dkhaii/cofeeshop-be/domain/product"
	prodmem "github.com/dkhaii/cofeeshop-be/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers costumer.CostumerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// loop througgh all the configs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// applies a customer repository to the OrderServive
func WithCustomerRepository(cr costumer.CostumerRepository) OrderConfiguration {
	// return a function that matches the order configuration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func (os *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) error {
	// fetch the customer
	cust, err := os.customers.Get(customerID)
	if err != nil {
		return err
	}

	// get each product
	var products []aggregate.Product
	var totalPrice float64

	for _, id := range productsIDs {
		p, err := os.products.GetByID(id)
		if err != nil {
			return err
		}

		products = append(products, p)
		totalPrice += p.GetPrice()
	}
	log.Printf("Costumer: %s has ordered %d product", cust.GetID(), len(products))

	return nil
}
