package businesslogic

import (
	"context"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"
	"tourmate/payment-service/model/entity"
)

type IPlatformFeedbackService interface {
	GetPlatformFeedbacks(req request.GetPlatformFeedbacksRequest, ctx context.Context) (response.PaginationDataResponse, error)
	GetPlatformFeedbackById(id int, ctx context.Context) (*entity.PlatformFeedback, error)
	CreatePlatformFeedback(req request.CreatePlatformFeedbackRequest, ctx context.Context) error
	UpdatePlatformFeedback(req request.UpdatePlatformFeedbackRequest, ctx context.Context) error
}
