package order_requests

type OrderShipmentRequest struct {
	id int
}

func (osr OrderShipmentRequest) GetId() int {
	return osr.id
}

func (osr *OrderShipmentRequest) SetId(id int) {
	osr.id = id
}
