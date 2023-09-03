package my_ent

import (
	"context"
	"github.com/yaroslavvasilenko/go_test_orm/orm/ent"
)

func (e *ClientEnt) CreateOrderForUser(order *ent.Order) (*ent.Order, error) {
	orderNew, err := e.Ent.Order.Create().
		SetName(order.Name).
		SetAmount(22).
		SetPrice(10).SetUserOrdersID(order.UserOrdersID).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return orderNew, nil

}
