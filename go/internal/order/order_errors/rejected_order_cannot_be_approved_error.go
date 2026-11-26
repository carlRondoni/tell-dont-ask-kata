package order_errors

type RejectedOrderCannotBeApproved struct {
	msg string
}

func (e RejectedOrderCannotBeApproved) Error() string {
	return e.msg
}

func NewRejectedOrderCannotBechangedError() error {
	return &RejectedOrderCannotBeApproved{
		msg: "Rejected order cannot be approved",
	}
}
