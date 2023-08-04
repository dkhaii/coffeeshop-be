package order

import (
	"testing"

	"github.com/dkhaii/cofeeshop-be/domain/customer"
	"github.com/dkhaii/cofeeshop-be/domain/product"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	esTeh, err := product.NewProduct("Es Teh", 5000, 2)
	if err != nil {
		t.Fatal(err)
	}

	cappucino, err := product.NewProduct("Cappucino", 10000, 2)
	if err != nil {
		t.Fatal(err)
	}

	zuppaSoup, err := product.NewProduct("Zuppa Soup", 25000, 1)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{
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

	cust, err := customer.NewCustomer("Mordekhai")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
