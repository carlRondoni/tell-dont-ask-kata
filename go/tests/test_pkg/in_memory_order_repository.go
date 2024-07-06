package test_pkg

import (
	"errors"
	"sync"

	"github.com/carlRondoni/tell-dont-ask-kata/internal/order"
)

type InMemoryOrderRepository struct {
	orders        []*order.Order
	insertedOrder *order.Order
	mu            sync.Mutex
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make([]*order.Order, 0),
	}
}

func (r *InMemoryOrderRepository) GetSavedOrder() *order.Order {
	return r.insertedOrder
}

func (r *InMemoryOrderRepository) Save(order order.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders = append(r.orders, &order)
	r.insertedOrder = &order
	return nil
}

func (r *InMemoryOrderRepository) AddOrder(order *order.Order) {
	r.Save(*order)
}

func (r *InMemoryOrderRepository) GetById(orderID int) (*order.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, o := range r.orders {
		if o.GetOrderId() == orderID {
			return o, nil
		}
	}
	return nil, errors.New("order not found")
}
