package request

type GetFeedbacksRequest struct {
	Request     SearchPaginationRequest `json:"request"`
	ServiceId   *int                    `json:"serviceId" form:"serviceId" binding:"omitempty,gt=0"`
	CustomerId  *int                    `json:"customerId" form:"customerId" binding:"omitempty,gt=0"`
	TourGuideId *int                    `json:"tourGuideId" form:"tourGuideId" binding:"omitempty,gt=0"`
	InvoiceId   *int                    `json:"invoiceId" form:"invoiceId" binding:"omitempty,gt=0"`
	Rating      *int                    `json:"rating" form:"rating" binding:"omitempty,gt=0"`
	IsDeleted   *bool                   `json:"isDeleted" form:"isDeleted" binding:"omitempty"`
	PageSize    int
}

type GetTourGuideFeedbacksRequest struct {
	PageSize    int `json:"pageSize" form:"pageSize" binding:"required,gt=0"`
	PageIndex   int `json:"pageIndex" form:"pageIndex"`
	TourGuideId int
}

type CreateFeedbackRequest struct {
	CustomerId  int    `json:"customerId" binding:"required,gt=0"`
	ServiceId   int    `json:"serviceId" binding:"required,gt=0"`
	TourGuideId int    `json:"tourGuideId" binding:"required,gt=0"`
	InvoiceId   int    `json:"invoiceId" binding:"required,gt=0"`
	Content     string `json:"content" binding:"required"`
	Rating      int    `json:"rating" binding:"required,gt=0,max=5"`
}

type RemoveFeedbackRequest struct {
	FeedbackId int `json:"feedbackId" binding:"required,gt=0"`
	ActorId    int `json:"actorId" binding:"required,gt=0"`
}

type UpdateFeedbackRequest struct {
	Request RemoveFeedbackRequest `json:"request"`
	Content string                `json:"content"`
	Rating  *int                  `json:"rating" binding:"omitempty,gt=0,max=5"`
}
