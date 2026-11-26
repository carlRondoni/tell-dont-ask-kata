package order

type Order struct {
	id       int
	status   string
	tax      float64
	total    float64
	items    []OrderItem
	currency string
}

func NewOrder(
	orderId int,
	status string,
) Order {
	return Order{
		id:     orderId,
		status: status,
		tax:    float64(0),
		total:  float64(0),
		items:  make([]OrderItem, 0),
	}
}

func (o Order) GetOrderId() int {
	return o.id
}

func (o Order) GetStatus() string {
	return o.status
}

func (o Order) GetTax() float64 {
	return o.tax
}

func (o Order) GetTotal() float64 {
	return o.total
}

func (o Order) GetItems() []OrderItem {
	return o.items
}

func (o Order) GetCurrency() string {
	return o.currency
}

func (o *Order) SetOrderId(orderId int) {
	o.id = orderId
}

func (o *Order) SetStatus(status string) {
	o.status = status
}

func (o *Order) SetTax(tax float64) {
	o.tax = tax
}

func (o *Order) SetTotal(total float64) {
	o.total = total
}

func (o *Order) SetItems(items []OrderItem) {
	o.items = items
}

func (o *Order) SetCurrency(currency string) {
	o.currency = currency
}
