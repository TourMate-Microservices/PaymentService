package entity

import "time"

type Payment struct {
	PaymentId     int       `json:"payment_id"`
	CustomerId    int       `json:"customer_id"`
	InvoiceId     int       `json:"invoice_id"`
	Price         float64   `json:"price"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
}

func (p Payment) GetPaymentTable() string {
	return "payments"
}
