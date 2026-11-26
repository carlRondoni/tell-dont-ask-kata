package order_cases

import (
	"math/rand"

	"github.com/carlRondoni/tell-dont-ask-kata/internal/order"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_errors"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_repository"
	"github.com/carlRondoni/tell-dont-ask-kata/internal/order/order_requests"
	catalog "github.com/carlRondoni/tell-dont-ask-kata/internal/product/product_catalog"
)

type OrderCreationUseCase struct {
	repository order_repository.OrderRepository
	catalog    catalog.ProductCatalog
}

func NewOrderCreationUseCase(repository order_repository.OrderRepository, catalog catalog.ProductCatalog) *OrderCreationUseCase {
	return &OrderCreationUseCase{
		repository: repository,
		catalog:    catalog,
	}
}

func (u *OrderCreationUseCase) Run(request *order_requests.SellItemsRequest) error {
	ord := order.NewOrder(
		rand.Intn(1000),
		order.Created.String(),
	)

	ord.SetCurrency("EUR")

	for _, itemRequest := range request.GetItems() {
		prod, err := u.catalog.GetByName(itemRequest.GetProductName())
		if err != nil {
			return err
		}
		if prod == nil {
			return order_errors.NewUnknownProductError()
		}

		unitaryTax := round((prod.GetPrice() / 100) * prod.GetCategory().GetTaxPercentage())
		unitaryTaxedAmount := round(prod.GetPrice() + unitaryTax)
		taxedAmount := round(unitaryTaxedAmount * float64(itemRequest.GetQuantity()))
		taxAmount := round(unitaryTax * float64(itemRequest.GetQuantity()))

		orderItem := order.OrderItem{}
		orderItem.SetQuantity(itemRequest.GetQuantity())
		orderItem.SetTax(taxAmount)
		orderItem.SetTaxedAmount(taxedAmount)
		orderItem.SetProduct(*prod)

		ord.SetItems(append(ord.GetItems(), orderItem))
		ord.SetTotal(ord.GetTotal() + taxedAmount)
		ord.SetTax(ord.GetTax() + taxAmount)

		err = u.repository.Save(ord)
		if err != nil {
			return err
		}
	}

	return nil
}

func round(amount float64) float64 {
	return float64(int((amount+0.005)*100)) / 100
}
