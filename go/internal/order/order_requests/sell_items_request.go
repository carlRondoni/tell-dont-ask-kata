package order_requests

type SellItemsRequest struct {
	items []SellItemRequest
}

func (sir SellItemsRequest) GetItems() []SellItemRequest {
	return sir.items
}

func (sir *SellItemsRequest) SetItems(items []SellItemRequest) {
	sir.items = items
}
