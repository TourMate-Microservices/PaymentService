package entity

import "time"

type Feedback struct {
	FeedbackId  int       `json:"feedbackId"`
	CustomerId  int       `json:"customerId"`
	TourGuideId int       `json:"tourGuideId"`
	CreatedDate time.Time `json:"createdDate"`
	Content     string    `json:"content"`
	Rating      int       `json:"rating"`
	IsDeleted   bool      `json:"isDeleted"`
	UpdatedAt   time.Time `json:"updatedAt"`
	InvoiceId   int       `json:"invoiceId"`
	ServiceId   int       `json:"serviceId"`
}

func (f Feedback) GetFeedbackTable() string {
	return "Feedback"
}

func (f Feedback) GetFeedbackLimitRecords() int {
	return 10
}
