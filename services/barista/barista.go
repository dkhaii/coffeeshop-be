package barista

import (
	"log"

	"github.com/dkhaii/cofeeshop-be/services/order"
	"github.com/google/uuid"
)

type BaristaConfiguration func(bar *Barista) error

type Barista struct {
	OrderService   *order.OrderService
	BillingService interface{}
}

func NewBarista(cfgs ...BaristaConfiguration) (*Barista, error) {
	bar := &Barista{}

	for _, cfg := range cfgs {
		err := cfg(bar)
		if err != nil {
			return nil, err
		}
	}

	return bar, nil
}

func WithOrderService(os *order.OrderService) BaristaConfiguration {
	return func(bar *Barista) error {
		bar.OrderService = os
		return nil
	}
}

func (bar *Barista) Order(customer uuid.UUID, products []uuid.UUID) error {
	totalPrice, err := bar.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("bill the costumer: %0.0f", totalPrice)

	return nil
}
