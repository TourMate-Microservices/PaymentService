package response

import "time"

type FeedbackResponse struct {
	FeedbackId  int       `json:"feedbackId"`
	CustomerId  int       `json:"customerId"`
	FullName    string    `json:"fullName"`
	Image       string    `json:"image"`
	Rating      int       `json:"rating"`
	Content     string    `json:"content"`
	CreatedDate time.Time `json:"createdDate"`
	ServiceId   int       `json:"serviceId"`
	ServiceName string    `json:"serviceName"`
}
