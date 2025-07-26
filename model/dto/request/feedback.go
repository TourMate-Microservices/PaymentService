package request

type GetFeedbacksRequest struct {
	Request     SearchPaginationRequest `json:"request"`
	ServiceId   *int                    `json:"serviceId" form:"serviceId" validate:"min=1"`
	CustomerId  *int                    `json:"customerId" form:"customerId" validate:"min=1"`
	TourGuideId *int                    `json:"tourGuideId" form:"tourGuideId" validate:"min=1"`
	InvoiceId   *int                    `json:"invoiceId" form:"invoiceId" validate:"min=1"`
	Rating      *int                    `json:"rating" form:"rating" validate:"min=1"`
	IsDeleted   *bool                   `json:"isDeleted" form:"isDeleted"`
}

type CreateFeedbackRequest struct {
	CustomerId  int    `json:"customerId" validate:"required, min=1"`
	ServiceId   int    `json:"serviceId" validate:"required, min=1"`
	TourGuideId int    `json:"tourGuideId" validate:"required, min=1"`
	InvoiceId   int    `json:"invoiceId" validate:"required, min=1"`
	Content     string `json:"content" validate:"required"`
	Rating      int    `json:"rating" validate:"required, min=1, max=5"`
}

type RemoveFeedbackRequest struct {
	FeedbackId int `json:"feedbackId" validate:"required, min=1"`
	ActorId    int `json:"actorId" validate:"required, min=1"`
}

type UpdateFeedbackRequest struct {
	Request RemoveFeedbackRequest `json:"request"`
	Content string                `json:"content"`
	Rating  *int                  `json:"rating"`
}
