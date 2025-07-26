package entity

import "time"

type Payment struct {
	//PaymentType   string    `json:"paymentType"` // bo^' bo? theo y' m
	PaymentId     int       `json:"paymentId"`
	Price         float64   `json:"price"`
	CreatedAt     time.Time `json:"createdAt"`
	PaymentMethod string    `json:"paymentMethod"`
	AccountId     int       `json:"accountId"`
	CustomerId    int       `json:"customerId"`
}

func (p Payment) GetPaymentTable() string {
	return "Payment"
}
