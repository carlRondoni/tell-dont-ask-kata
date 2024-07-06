package order_cases_test

import (
	"testing"

	"github.com/carlRondoni/tell-dont-ask-kata/internal/order"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_cases"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_requests"
	"github.com/carlRondoni/tell-dont-ask-kata/tests/test_pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setUpOrderShipmentUseCaseTest() (*order_cases.OrderShipmentUseCase, *test_pkg.InMemoryOrderRepository, *test_pkg.InMemoryShipmentService) {
	orderRepository := test_pkg.NewInMemoryOrderRepository()
	shipmentService := test_pkg.NewInMemoryShipmentService()
	useCase := order_cases.NewOrderShipmentUseCase(orderRepository, shipmentService)

	return useCase, orderRepository, shipmentService
}

func TestShipApprovedOrder(t *testing.T) {
	useCase, orderRepository, _ := setUpOrderShipmentUseCaseTest()

	initialOrder := order.NewOrder(
		1,
		order.Approved.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := &order_requests.OrderShipmentRequest{}
	request.SetId(1)

	err := useCase.Run(request)
	require.NoError(t, err)

	savedOrder := orderRepository.GetSavedOrder()
	assert.Equal(t, order.Shipped.String(), savedOrder.GetStatus())
}

func TestCreatedOrdersCannotBeShipped(t *testing.T) {
	useCase, orderRepository, shipmentService := setUpOrderShipmentUseCaseTest()

	initialOrder := order.NewOrder(
		1,
		order.Created.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := &order_requests.OrderShipmentRequest{}
	request.SetId(1)

	err := useCase.Run(request)
	require.Error(t, err)
	assert.Equal(t, "order cannot be shipped", err.Error())

	assert.Nil(t, orderRepository.GetSavedOrder())
	assert.Nil(t, shipmentService.GetShippedOrder())
}

func TestRejectedOrdersCannotBeShipped(t *testing.T) {
	useCase, orderRepository, shipmentService := setUpOrderShipmentUseCaseTest()

	initialOrder := order.NewOrder(
		1,
		order.Rejected.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := &order_requests.OrderShipmentRequest{}
	request.SetId(1)

	err := useCase.Run(request)
	require.Error(t, err)
	assert.Equal(t, "order cannot be shipped", err.Error())

	assert.Nil(t, orderRepository.GetSavedOrder())
	assert.Nil(t, shipmentService.GetShippedOrder())
}

func TestShippedOrdersCannotBeShippedAgain(t *testing.T) {
	useCase, orderRepository, shipmentService := setUpOrderShipmentUseCaseTest()

	initialOrder := order.NewOrder(
		1,
		order.Shipped.String(),
	)
	orderRepository.AddOrder(&initialOrder)

	request := &order_requests.OrderShipmentRequest{}
	request.SetId(1)

	err := useCase.Run(request)
	require.Error(t, err)
	assert.Equal(t, "order cannot be shipped twice", err.Error())

	assert.Nil(t, orderRepository.GetSavedOrder())
	assert.Nil(t, shipmentService.GetShippedOrder())
}
