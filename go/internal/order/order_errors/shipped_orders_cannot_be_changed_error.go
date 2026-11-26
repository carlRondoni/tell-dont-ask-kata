package order_errors

type ShippedOrdersCannotBeChanged struct {
	msg string
}

func (e ShippedOrdersCannotBeChanged) Error() string {
	return e.msg
}

func NewShippedOrdersCannotBechangedError() error {
	return &ShippedOrdersCannotBeChanged{
		msg: "Shipped orders cannot be changed",
	}
}
