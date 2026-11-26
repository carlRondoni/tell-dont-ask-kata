package order_requests

type SellItemRequest struct {
	productName string
	quantity    int
}

func (sir SellItemRequest) GetProductName() string {
	return sir.productName
}

func (sir *SellItemRequest) SetProductName(productName string) {
	sir.productName = productName
}

func (sir SellItemRequest) GetQuantity() int {
	return sir.quantity
}

func (sir *SellItemRequest) SetQuantity(quantity int) {
	sir.quantity = quantity
}
