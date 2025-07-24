package request

type GetPaymentsRequest struct {
	Request    SearchPaginationRequest `json:"request"`
	Method     string                  `json:"method" form:"method"`
	Status     string                  `json:"status" form:"status"`
	CustomerId *int                    `json:"customer_id" form:"customer_id"`
}

type CreatePaymentRequest struct {
	CustomerId    int     `json:"customer_id" validate:"required, min=1"`
	InvoiceId     int     `json:"invoice_id" validate:"required, min=1"`
	Price         float64 `json:"price" validate:"required, min=1"`
	PaymentMethod string  `json:"payment_method" validate:"required"`
}

type UpdatePaymentRequest struct {
	PaymentId int    `json:"payment_id" validate:"required"`
	Method    string `json:"method"`
	Status    string `json:"status"`
}
