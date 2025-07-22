package repo

import (
	"context"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/entity"
)

type IFeedbackRepo interface {
	// Response data: data, total pages, total records, error
	GetFeedbacks(req request.GetFeedbacksRequest, ctx context.Context) (*[]entity.Feedback, int, int, error)
	GetFeedbackById(id int, ctx context.Context) (*entity.Feedback, error)
	CreateFeedback(feedback entity.Feedback, ctx context.Context) error
	UpdateFeedback(feedback entity.Feedback, ctx context.Context) error
}
