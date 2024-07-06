package order

type OrderStatus int

const (
	Created OrderStatus = iota
	Approved
	Shipped
	Rejected
)

func (status OrderStatus) String() string {
	return [...]string{"Created", "Approved", "Shipped", "Rejected"}[status]
}
