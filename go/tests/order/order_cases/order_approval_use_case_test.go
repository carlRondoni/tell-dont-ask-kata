package order_cases_test

import (
	"testing"

	"github.com/carlRondoni/tell-dont-ask-kata/internal/order"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_cases"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_errors"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_requests"
	"github.com/carlRondoni/tell-dont-ask-kata/tests/test_pkg"
	"github.com/stretchr/testify/assert"
)

func TestOrderApprovalUseCase_ApproveExistingOrder(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Created,
	)
	orderRepository.AddOrder(&initialOrder)

	request := &order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(true)
	err := useCase.Run(request)

	assert.NoError(t, err)
	savedOrder := orderRepository.GetSavedOrder()
	assert.Equal(t, order.Approved, savedOrder.GetStatus())
}

func TestOrderApprovalUseCase_RejectExistingOrder(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Created,
	)
	orderRepository.AddOrder(&initialOrder)

	request := order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(false)
	err := useCase.Run(&request)

	assert.NoError(t, err)
	savedOrder := orderRepository.GetSavedOrder()
	assert.Equal(t, order.Rejected, savedOrder.GetStatus())
}

func TestOrderApprovalUseCase_CannotApproveRejectedOrder(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Rejected,
	)
	orderRepository.AddOrder(&initialOrder)

	request := order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(true)
	err := useCase.Run(&request)

	var target *order_errors.RejectedOrderCannotBeApproved
	assert.ErrorAs(t, err, &target)
	assert.Nil(t, orderRepository.GetSavedOrder())
}

func TestOrderApprovalUseCase_CannotRejectApprovedOrder(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Approved,
	)
	orderRepository.AddOrder(&initialOrder)

	request := order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(false)
	err := useCase.Run(&request)

	var target *order_errors.ApprovedOrderCannotBeRejected
	assert.ErrorAs(t, err, &target)
	assert.Nil(t, orderRepository.GetSavedOrder())
}

func TestOrderApprovalUseCase_ShippedOrdersCannotBeRejected(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Shipped,
	)
	orderRepository.AddOrder(&initialOrder)

	request := order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(false)
	err := useCase.Run(&request)

	var target *order_errors.ShippedOrdersCannotBeChanged
	assert.ErrorAs(t, err, &target)
	assert.Nil(t, orderRepository.GetSavedOrder())
}
