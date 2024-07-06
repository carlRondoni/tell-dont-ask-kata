package order_errors

type UnknownProduct struct {
	msg string
}

func (e UnknownProduct) Error() string {
	return e.msg
}

func NewUnknownProductError() error {
	return &UnknownProduct{
		msg: "unknown product",
	}
}
