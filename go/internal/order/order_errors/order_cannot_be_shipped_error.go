package order_errors

type OrderCannotBeShipped struct {
	msg string
}

func (e OrderCannotBeShipped) Error() string {
	return e.msg
}

func NewOrderCannotBeShippedError() error {
	return &OrderCannotBeShipped{
		msg: "Order cannot be shipped",
	}
}
