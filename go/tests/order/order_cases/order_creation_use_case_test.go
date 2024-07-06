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

func setUpOrderCreationUseCaseTest() (*order_cases.OrderCreationUseCase, *test_pkg.InMemoryOrderRepository) {
	productCatalog := test_pkg.NewInMemoryProductCatalog()

	orderRepository := test_pkg.NewInMemoryOrderRepository()
	orderCase := order_cases.NewOrderCreationUseCase(
		orderRepository,
		productCatalog,
	)

	return orderCase, orderRepository
}

func TestMultipleItems(t *testing.T) {
	useCase, orderRepository := setUpOrderCreationUseCaseTest()

	saladRequest := order_requests.SellItemRequest{}
	saladRequest.SetProductName("salad")
	saladRequest.SetQuantity(2)
	tomatoRequest := order_requests.SellItemRequest{}
	tomatoRequest.SetProductName("tomato")
	tomatoRequest.SetQuantity(3)

	request := order_requests.SellItemsRequest{}
	request.SetItems(append(request.GetItems(), saladRequest, tomatoRequest))

	err := useCase.Run(&request)
	require.NoError(t, err)

	insertedOrder := orderRepository.GetSavedOrder()
	assert.Equal(t, order.Created.String(), insertedOrder.GetStatus())
	assert.Equal(t, 23.20, insertedOrder.GetTotal())
	assert.Equal(t, 2.13, insertedOrder.GetTax())
	assert.Equal(t, "EUR", insertedOrder.GetCurrency())
	assert.Len(t, insertedOrder.GetItems(), 2)
	assert.Equal(t, "salad", insertedOrder.GetItems()[0].GetProduct().GetName())
	assert.Equal(t, 3.56, insertedOrder.GetItems()[0].GetProduct().GetPrice())
	assert.Equal(t, 2, insertedOrder.GetItems()[0].GetQuantity())
	assert.Equal(t, 7.84, insertedOrder.GetItems()[0].GetTaxedAmount())
	assert.Equal(t, 0.72, insertedOrder.GetItems()[0].GetTax())
	assert.Equal(t, "tomato", insertedOrder.GetItems()[1].GetProduct().GetName())
	assert.Equal(t, 4.65, insertedOrder.GetItems()[1].GetProduct().GetPrice())
	assert.Equal(t, 3, insertedOrder.GetItems()[1].GetQuantity())
	assert.Equal(t, 15.36, insertedOrder.GetItems()[1].GetTaxedAmount())
	assert.Equal(t, 1.41, insertedOrder.GetItems()[1].GetTax())
}

func TestUnknownProduct(t *testing.T) {
	useCase, _ := setUpOrderCreationUseCaseTest()

	unknown := order_requests.SellItemRequest{}
	unknown.SetProductName("unknown product")

	request := order_requests.SellItemsRequest{}
	request.SetItems(append(request.GetItems(), unknown))

	err := useCase.Run(&request)
	require.Error(t, err)
	assert.EqualError(t, err, "unknown product")
}
