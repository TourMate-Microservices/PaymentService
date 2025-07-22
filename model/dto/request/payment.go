package request

type GetPaymentsRequest struct {
	Request    SearchPaginationRequest `json:"request"`
	Method     string                  `json:"method" form:"method"`
	Status     string                  `json:"status" form:"status"`
	CustomerId *int                    `json:"customer_id" form:"customer_id"`
}

type UpdatePaymentRequest struct {
	PaymentId int    `json:"payment_id" validate:"required"`
	Method    string `json:"method"`
	Status    string `json:"status"`
}
