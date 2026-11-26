package order_cases

import (
	"errors"

	"github.com/carlRondoni/tell-dont-ask-kata/internal/order"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_errors"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_repository"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_requests"
)

type OrderShipmentUseCase struct {
	repository      order_repository.OrderRepository
	shipmentService order.ShipmentService
}

func NewOrderShipmentUseCase(repository order_repository.OrderRepository, shipmentService order.ShipmentService) *OrderShipmentUseCase {
	return &OrderShipmentUseCase{
		repository:      repository,
		shipmentService: shipmentService,
	}
}

func (u *OrderShipmentUseCase) Run(request *order_requests.OrderShipmentRequest) error {
	ord, err := u.repository.GetById(request.GetId())
	if err != nil {
		return err
	}
	if ord == nil {
		return errors.New("Order not found")
	}

	switch ord.GetStatus() {
	case order.Created.String(), order.Rejected.String():
		return order_errors.NewOrderCannotBeShippedError()
	case order.Shipped.String():
		return order_errors.NewOrderCannotBeShippedTwiceError()

	}

	u.shipmentService.Ship(ord)

	ord.SetStatus(order.Shipped.String())
	u.repository.Save(*ord)

	return nil
}
