package request

type GetPaymentsRequest struct {
	Request    SearchPaginationRequest `json:"request"`
	Method     string                  `json:"method" form:"method"`
	CustomerId *int                    `json:"customerId" form:"customerId"`
}

type CreatePaymentRequest struct {
	CustomerId    int     `json:"customerId" validate:"required, min=1"`
	AccountId     int     `json:"accountId" validate:"required, min=1"`
	Price         float64 `json:"price" validate:"required, min=1"`
	PaymentType   string  `json:"paymentType" validate:"required"`
	PaymentMethod string  `json:"paymentMethod" validate:"required"`
}

type UpdatePaymentRequest struct {
	PaymentId int    `json:"paymentId" validate:"required"`
	Method    string `json:"method"`
}
