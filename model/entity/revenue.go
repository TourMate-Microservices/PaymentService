package entity

import "time"

type Revenue struct {
	RevenueId          int       `json:"revenueId"`
	PaymentId          int       `json:"paymentId"`
	TourGuideId        int       `json:"tourGuideId"`
	InvoiceId          int       `json:"invoiceId"`
	TotalAmount        float64   `json:"totalAmount"`
	ActualReceived     float64   `json:"actualReceived"`
	PlatformCommission float64   `json:"platformCommission"`
	PaymentStatus      bool      `json:"paymentStatus"`
	CreatedAt          time.Time `json:"createdAt"`
}

func (r Revenue) GetRevenueTable() string {
	return "Revenue"
}
