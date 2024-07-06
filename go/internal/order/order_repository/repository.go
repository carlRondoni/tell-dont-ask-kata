package order_repository

import (
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order"
)

type OrderRepository interface {
	Save(order order.Order) error
	GetById(orderId int) (*order.Order, error)
}
