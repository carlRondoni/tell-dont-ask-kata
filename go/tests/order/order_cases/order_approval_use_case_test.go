package order_cases_test

import (
	"testing"

	"github.com/carlRondoni/tell-dont-ask-kata/internal/order"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_cases"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_requests"
	"github.com/carlRondoni/tell-dont-ask-kata/tests/test_pkg"
	"github.com/stretchr/testify/assert"
)

func TestOrderApprovalUseCase_ApproveExistingOrder(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Created.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := &order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(true)
	err := useCase.Run(request)

	assert.NoError(t, err)
	savedOrder := orderRepository.GetSavedOrder()
	assert.Equal(t, order.Approved.String(), savedOrder.GetStatus())
}

func TestOrderApprovalUseCase_RejectExistingOrder(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Created.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(false)
	err := useCase.Run(&request)

	assert.NoError(t, err)
	savedOrder := orderRepository.GetSavedOrder()
	assert.Equal(t, order.Rejected.String(), savedOrder.GetStatus())
}

func TestOrderApprovalUseCase_CannotApproveRejectedOrder(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Rejected.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(true)
	err := useCase.Run(&request)

	assert.Error(t, err, "should throw RejectedOrderCannotBeApprovedException")
	assert.Nil(t, orderRepository.GetSavedOrder())
}

func TestOrderApprovalUseCase_CannotRejectApprovedOrder(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Approved.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(true)
	err := useCase.Run(&request)

	assert.Error(t, err, "should throw ApprovedOrderCannotBeRejectedException")
	assert.Nil(t, orderRepository.GetSavedOrder())
}

func TestOrderApprovalUseCase_ShippedOrdersCannotBeRejected(t *testing.T) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	useCase := order_cases.NewOrderApprovalUseCase(orderRepository)

	initialOrder := order.NewOrder(
		1,
		order.Shipped.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := order_requests.OrderApprovalRequest{}
	request.SetId(1)
	request.SetApproved(false)
	err := useCase.Run(&request)

	assert.Error(t, err, "should throw ShippedOrdersCannotBeChangedException")
	assert.Nil(t, orderRepository.GetSavedOrder())
}
