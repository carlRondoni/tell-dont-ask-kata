package test_pkg

import "github.com/carlRondoni/tell-dont-ask-kata/internal/order"

type InMemoryShipmentService struct {
	shippedOrder *order.Order
}

func NewInMemoryShipmentService() *InMemoryShipmentService {
	return &InMemoryShipmentService{}
}

func (s InMemoryShipmentService) GetShippedOrder() *order.Order {
	return s.shippedOrder
}

func (s *InMemoryShipmentService) Ship(order *order.Order) {
	s.shippedOrder = order
}
