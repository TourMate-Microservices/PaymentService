package entity

import "time"

type Revenue struct {
	RevenueId          int       `json:"revenue_id"`
	TourGuideId        int       `json:"tour_guide_id"`
	InvoiceId          int       `json:"invoice_id"`
	TotalAmount        float64   `json:"total_amount"`
	ActualReceived     float64   `json:"acutal_recieved"`
	PlatformCommission float64   `json:"platform_commision"`
	PaymentStatus      bool      `json:"payment_status"`
	CreatedAt          time.Time `json:"created_at"`
}

func (r Revenue) GetRevenueTable() string {
	return "Revenue"
}
