package entity

import "time"

type PlatformFeedback struct {
	FeedbackId int       `json:"feedback_id"`
	AccountId  int       `json:"account_id"`
	PaymentId  int       `json:"payment_id"`
	Content    string    `json:"content"`
	Rating     int       `json:"rating"`
	CreatedAt  time.Time `json:"created_at"`
}

func (p PlatformFeedback) GetPlatformFeedbackTable() string {
	return "PlatformFeedback"
}

func (p PlatformFeedback) GetPlatformFeedbackLimitRecords() int {
	return 10
}
