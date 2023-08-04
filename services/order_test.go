package services

import (
	"testing"

	"github.com/dkhaii/cofeeshop-be/aggregate"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	esTeh, err := aggregate.NewProduct("Es Teh", 5000, 2)
	if err != nil {
		t.Fatal(err)
	}

	cappucino, err := aggregate.NewProduct("Cappucino", 10000, 2)
	if err != nil {
		t.Fatal(err)
	}

	zuppaSoup, err := aggregate.NewProduct("Zuppa Soup", 25000, 1)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		esTeh,
		cappucino,
		zuppaSoup,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Mordekhai")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err2 := os.CreateOrder(cust.GetID(), order)
	if err2 != nil {
		t.Error(err2)
	}
}
