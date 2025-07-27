package entity

import "time"

type Payment struct {
	//PaymentType   string    `json:"paymentType"` // bo^' bo? theo y' m
	PaymentId     int       `json:"paymentId"`
	Price         float64   `json:"price"`
	CreatedAt     time.Time `json:"createdAt"`
	PaymentMethod string    `json:"paymentMethod"`
	InvoiceId     int       `json:"invoiceId"`
	CustomerId    int       `json:"customerId"`
	ServiceId     int       `json:"serviceId"`
	Status        string    `json:"status"` // e.g., "paid", "unpaid", "pending"
}

func (p Payment) GetPaymentTable() string {
	return "Payment"
}
