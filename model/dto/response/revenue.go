package response

import "time"

type MonthlyRevenueResponse struct {
	Month             int     `json:"month"`
	Year              int     `json:"year"`
	TotalRevenue      float64 `json:"totalRevenue"`
	PlatformFee       float64 `json:"platformFee"`
	NetRevenue        float64 `json:"netRevenue"`
	TotalRecords      int     `json:"totalRecords"`
	CompletedPayments int     `json:"completedPayments"`
	PendingPayments   int     `json:"pendingPayments"`
	GrowthPercentage  float64 `json:"growthPercentage"`
}

type RevenueResponse struct {
	RevenueId          int       `json:"revenueId"`
	PaymentId          int       `json:"paymentId"`
	InvoiceId          int       `json:"invoiceId"`
	TourGuideId        int       `json:"tourGuideId"`
	TotalAmount        float64   `json:"totalAmount"`
	ActualReceived     float64   `json:"actualReceived"`
	PlatformCommission float64   `json:"platformCommission"`
	CreatedAt          time.Time `json:"createdAt"`
	PaymentStatus      bool      `json:"paymentStatus"`
	TourGuideName      string    `json:"tourGuideName"`
}

type RevenueGrowthPercentageResponse struct {
	GrowthPercentage float64 `json:"growthPercentage "`
}

type RevenueStatusResponse struct {
	TotalRevenue      float64           `json:"totalRevenue"`
	PlatformFee       float64           `json:"platformFee"`
	NetRevenue        float64           `json:"netRevenue"`
	TotalRecords      int               `json:"totalRecords"`
	CompletedPayments int               `json:"completedPayments"`
	PendingPayments   int               `json:"pendingPayments"`
	MonthlyGrowth     float64           `json:"monthlyGrowth"`
	RevenueList       []RevenueResponse `json:"revenueList"`
}
