package order_errors

type OrderCannotBeShippedTwice struct {
	msg string
}

func (e OrderCannotBeShippedTwice) Error() string {
	return e.msg
}

func NewOrderCannotBeShippedTwiceError() error {
	return &OrderCannotBeShippedTwice{
		msg: "order cannot be shipped twice",
	}
}
