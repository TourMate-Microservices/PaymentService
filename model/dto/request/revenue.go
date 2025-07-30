package request

type GetRevenuesRequest struct {
	TourGuideId   int   `json:"tourGuideId" form:"tourGuideId" binding:"required,gt=0"`
	Year          *int  `json:"year" form:"year" binding:"omitempty,gt=2020"`
	Month         *int  `json:"month" form:"month" binding:"omitempty,gt=0,max=12"`
	PaymentStatus *bool `json:"paymentStatus" form:"paymentStatus"`
	PageNumber    *int  `json:"pageNumber" form:"pageNumber" binding:"omitempty,gt=0"`
	PageSize      *int  `json:"pageSize" form:"pageSize" binding:"omitempty,gt=0"`
}

type GetMonthlyRevenueRequest struct {
	TourGuideId int `json:"tourGuideId" form:"tourGuideId" binding:"required,gt=0"`
	Year        int `json:"year" form:"year" binding:"required,gt=2020"`
	Month       int `json:"month" form:"month" binding:"required,gt=0,max=12"`
}

type CreateRevenueRequest struct {
	PaymentId          int     `json:"paymentId" binding:"required,gt=0"`
	TourGuideId        int     `json:"tourGuideId" binding:"required,gt=0"`
	InvoiceId          int     `json:"invoiceId" binding:"required,gt=0"`
	TotalAmount        float64 `json:"totalAmount" binding:"required,gt=0"`
	ActualReceived     float64 `json:"actualReceived" binding:"required,gt=0"`
	PlatformCommission float64 `json:"platformCommission" binding:"required,gt=0"`
	PaymentStatus      bool    `json:"paymentStatus" binding:"required"`
}

type UpdateRevenueRequest struct {
	RevenueId          int      `json:"revenueId"`
	PaymentId          *int     `json:"paymentId" binding:"omitempty,gt=0"`
	TourGuideId        *int     `json:"tourGuideId" binding:"omitempty,gt=0"`
	InvoiceId          *int     `json:"invoiceId" binding:"omitempty,gt=0"`
	TotalAmount        *float64 `json:"totalAmount" binding:"omitempty,gt=0"`
	ActualReceived     *float64 `json:"actualReceived" binding:"omitempty,gt=0"`
	PlatformCommission *float64 `json:"platformCommission" binding:"omitempty,gt=0"`
	PaymentStatus      *bool    `json:"paymentStatus" binding:"omitempty"`
}
