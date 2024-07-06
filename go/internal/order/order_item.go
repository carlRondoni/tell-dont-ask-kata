package order

import "github.com/carlRondoni/tell-dont-ask-kata/internal/product"

type OrderItem struct {
	quantity    int
	tax         float64
	taxedAmount float64
	product     product.Product
}

func (oi OrderItem) GetQuantity() int {
	return oi.quantity
}

func (oi *OrderItem) SetQuantity(quantity int) {
	oi.quantity = quantity
}

func (oi OrderItem) GetTax() float64 {
	return oi.tax
}

func (oi *OrderItem) SetTax(tax float64) {
	oi.tax = tax
}

func (oi OrderItem) GetTaxedAmount() float64 {
	return oi.taxedAmount
}

func (oi *OrderItem) SetTaxedAmount(taxedAmount float64) {
	oi.taxedAmount = taxedAmount
}

func (oi OrderItem) GetProduct() product.Product {
	return oi.product
}

func (oi *OrderItem) SetProduct(prod product.Product) {
	oi.product = prod
}
