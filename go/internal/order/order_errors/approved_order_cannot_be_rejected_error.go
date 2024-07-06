package order_errors

type ApprovedOrderCannotBeRejected struct {
	msg string
}

func (e ApprovedOrderCannotBeRejected) Error() string {
	return e.msg
}

func NewApprovedOrderCannotBeRejectedError() error {
	return &ApprovedOrderCannotBeRejected{
		msg: "Approved order cannot be rejected",
	}
}
