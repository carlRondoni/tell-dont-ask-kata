package order

type ShipmentService interface {
	Ship(order *Order)
}
