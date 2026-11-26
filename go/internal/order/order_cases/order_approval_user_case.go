package order_cases

import (
	"errors"

	"github.com/carlRondoni/tell-dont-ask-kata/internal/order"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_errors"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_repository"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_requests"
)

type OrderApprovalUseCase struct {
	repository order_repository.OrderRepository
}

func NewOrderApprovalUseCase(repository order_repository.OrderRepository) *OrderApprovalUseCase {
	return &OrderApprovalUseCase{
		repository: repository,
	}
}

func (u *OrderApprovalUseCase) Run(request *order_requests.OrderApprovalRequest) error {
	ord, err := u.repository.GetById(request.GetId())
	if err != nil {
		return err
	}
	if ord == nil {
		return errors.New("order not found")
	}

	switch ord.GetStatus() {
	case order.Shipped:
		return order_errors.NewShippedOrdersCannotBechangedError()
	case order.Rejected:
		if request.IsApproved() {
			return order_errors.NewRejectedOrderCannotBechangedError()
		}
	case order.Approved:
		if !request.IsApproved() {
			return order_errors.NewApprovedOrderCannotBeRejectedError()
		}
	}

	var newStatus order.OrderStatus
	if request.IsApproved() {
		newStatus = order.Approved
	} else {
		newStatus = order.Rejected
	}

	ord.SetStatus(newStatus)
	err = u.repository.Save(*ord)
	if err != nil {
		return err
	}

	return nil
}
