package request

type GetPlatformFeedbacksRequest struct {
	Request   SearchPaginationRequest `json:"request"`
	AccountId *int                    `json:"account_id" form:"account_id" validate:"min=1"`
	Rating    *int                    `json:"rating" form:"rating" validate:"min=1"`
}

type CreatePlatformFeedbackRequest struct {
	AccountId int    `json:"account_id" validate:"required, min=1"`
	PaymentId int    `json:"payment_id" validate:"required, min=1"`
	Content   string `json:"content" validte:"required"`
	Rating    int    `json:"rating" validate:"required, min=1, max=5"`
}

type UpdatePlatformFeedbackRequest struct {
	FeedbackId int    `json:"feedback_id" validate:"required, min=1"`
	ActorId    int    `json:"actor_id" validate:"required, min=1"`
	Content    string `json:"content"`
	Rating     *int   `json:"rating"`
}
