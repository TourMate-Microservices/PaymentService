package entity

import "time"

type Feedback struct {
	FeedbackId  int       `json:"feedback_id"`
	CustomerId  int       `json:"customer_id"`
	TourGuideId int       `json:"tour_guide_id"`
	InvoiceId   int       `json:"invoice_id"`
	Content     string    `json:"content"`
	Rating      int       `json:"rating"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (f Feedback) GetFeedbackTable() string {
	return "Feedback"
}

func (f Feedback) GetFeedbackLimitRecords() int {
	return 10
}
