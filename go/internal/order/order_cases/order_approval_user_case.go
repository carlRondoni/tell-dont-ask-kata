package order_cases

import (
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

	switch ord.GetStatus() {
	case order.Shipped.String():
		return order_errors.NewShippedOrdersCannotBechangedError()
	case order.Rejected.String():
		if request.IsApproved() {
			return order_errors.NewRejectedOrderCannotBechangedError()
		}
	case order.Approved.String():
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

	ord.SetStatus(newStatus.String())
	err = u.repository.Save(*ord)
	if err != nil {
		return err
	}

	return nil
}
