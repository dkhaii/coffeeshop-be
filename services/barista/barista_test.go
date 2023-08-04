package barista

import (
	"testing"

	"github.com/dkhaii/cofeeshop-be/domain/product"
	"github.com/dkhaii/cofeeshop-be/services/order"
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

func Test_Barista(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	barista, err := NewBarista(
		WithOrderService(os),
	)
	if err != nil {
		t.Fatal(err)
	}

	custID, err := os.AddCustomer("Mordekhai")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[1].GetID(),
		products[2].GetID(),
	}

	err = barista.Order(custID, order)
	if err != nil {
		t.Fatal(err)
	}
}
