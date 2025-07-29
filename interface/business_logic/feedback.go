package businesslogic

import (
	"context"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"
	"tourmate/payment-service/model/entity"
)

type IFeedbackService interface {
	GetFeedbacks(req request.GetFeedbacksRequest, ctx context.Context) (response.PaginationDataResponse, error)
	GetTourGuideFeedbacks(req request.GetTourGuideFeedbacksRequest, ctx context.Context) (response.PaginationDataResponse, error)
	GetFeedbackById(id int, ctx context.Context) (*entity.Feedback, error)
	CreateFeedback(req request.CreateFeedbackRequest, ctx context.Context) (*entity.Feedback, error)
	UpdateFeedback(req request.UpdateFeedbackRequest, ctx context.Context) (*entity.Feedback, error)
	RemoveFeedback(req request.RemoveFeedbackRequest, ctx context.Context) error
}
