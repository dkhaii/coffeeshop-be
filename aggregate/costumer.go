package aggregate

import "github.com/dkhaii/cofeeshop-be/entity"

type Costumer struct {
	person *entity.Person
	product []*entity.Item
	transactions []*entity.Transaction
}