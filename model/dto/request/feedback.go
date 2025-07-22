package request

type GetFeedbacksRequest struct {
	Request     SearchPaginationRequest `json:"request"`
	CustomerId  *int                    `json:"customer_id" form:"customer_id" validate:"min=1"`
	TourGuideId *int                    `json:"tour_guide_id" form:"tour_guide_id" validate:"min=1"`
	InvoiceId   *int                    `json:"invoice_id" form:"invoice_id" validate:"min=1"`
	Rating      *int                    `json:"rating" form:"rating" validate:"min=1"`
	IsDeleted   *bool                   `json:"is_deleted" form:"is_deleted"`
}

type CreateFeedbackRequest struct {
	CustomerId  int    `json:"customer_id" validate:"required, min=1"`
	TourGuideId int    `json:"tour_guide_id" validate:"required, min=1"`
	InvoiceId   int    `json:"invoice_id" validate:"required, min=1"`
	Content     string `json:"content" validte:"required"`
	Rating      int    `json:"rating" validate:"required, min=1, max=5"`
}

type RemoveFeedbackRequest struct {
	FeedbackId int `json:"feedback_id" validate:"required, min=1"`
	ActorId    int `json:"actor_id" validate:"required, min=1"`
}

type UpdateFeedbackRequest struct {
	Request RemoveFeedbackRequest `json:"request"`
	Content string                `json:"content"`
	Rating  *int                  `json:"rating"`
}
