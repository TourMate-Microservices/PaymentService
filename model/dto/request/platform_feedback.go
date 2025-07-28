package request

type GetPlatformFeedbacksRequest struct {
	Request    SearchPaginationRequest `json:"request"`
	CustomerId *int                    `json:"customerId" form:"customerId" binding:"omitempty,gt=0"`
	Rating     *int                    `json:"rating" form:"rating" binding:"omitempty,gt=0,max=5"`
	PageIndex  *int                    `json:"pageIndex" form:"pageIndex" binding:"omitempty,gt=0"`
	PageSize   *int                    `json:"pageSize" form:"pageSize" binding:"omitempty,gt=0"`
}

type CreatePlatformFeedbackRequest struct {
	CustomerId int    `json:"customerId" binding:"required,gt=0"`
	PaymentId  int    `json:"paymentId" binding:"required,gt=0"`
	Content    string `json:"content" validte:"required"`
	Rating     int    `json:"rating" binding:"required,gt=0,max=5"`
}

type UpdatePlatformFeedbackRequest struct {
	FeedbackId int    `json:"feedbackId" binding:"required,gt=0"`
	ActorId    int    `json:"actorId" binding:"required,gt=0"`
	Content    string `json:"content"`
	Rating     *int   `json:"rating"`
}
