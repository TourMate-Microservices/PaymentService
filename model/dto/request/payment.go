package request

type GetPaymentsRequest struct {
	Request    SearchPaginationRequest `json:"request"`
	Method     string                  `json:"method" form:"method"`
	CustomerId *int                    `json:"customerId" form:"customerId" binding:"omitempty,gt=0"`
}

type CreatePaymentRequest struct {
	CustomerId    int     `json:"customerId" binding:"required,gt=0"`
	InvoiceId     int     `json:"invoiceId" binding:"required,gt=0"`
	ServiceId     int     `json:"serviceId" binding:"required,gt=0"`
	Price         float64 `json:"price" binding:"required,gt=0"`
	PaymentMethod string  `json:"paymentMethod" binding:"required"`
}

type UpdatePaymentRequest struct {
	PaymentId int    `json:"paymentId" binding:"required"`
	Method    string `json:"method"`
}

type CreatePayosTransactionRequest struct {
	Amount    float64 `json:"amount" binding:"required,gt=0"`
	InvoiceId int     `json:"invoiceId" binding:"required,gt=0"`
}
