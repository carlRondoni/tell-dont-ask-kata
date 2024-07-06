package order_requests

type OrderApprovalRequest struct {
	id       int
	approved bool
}

func NewOrderApprovalRequest() *OrderApprovalRequest {
	return &OrderApprovalRequest{}
}

func (o *OrderApprovalRequest) GetId() int {
	return o.id
}

func (o *OrderApprovalRequest) SetId(id int) *OrderApprovalRequest {
	o.id = id
	return o
}

func (o *OrderApprovalRequest) IsApproved() bool {
	return o.approved
}

func (o *OrderApprovalRequest) SetApproved(approved bool) *OrderApprovalRequest {
	o.approved = approved
	return o
}
