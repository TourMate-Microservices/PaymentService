package entity

import "time"

type PlatformFeedback struct {
	FeedbackId int       `json:"feedbackId"`
	CustomerId int       `json:"customerId"`
	PaymentId  int       `json:"paymentId"`
	Content    string    `json:"content"`
	Rating     int       `json:"rating"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (p PlatformFeedback) GetPlatformFeedbackTable() string {
	return "PlatformFeedback"
}

func (p PlatformFeedback) GetPlatformFeedbackLimitRecords() int {
	return 10
}
