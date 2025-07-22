package businesslogic

import (
	"context"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"

	"tourmate/payment-service/model/entity"
)

type IPaymentService interface {
	GetPayments(req request.GetPaymentsRequest, ctx context.Context) (response.PaginationDataResponse, error)
	GetPaymentById(id int, ctx context.Context) (*entity.Payment, error)
	UpdatePayment(req request.UpdatePaymentRequest, ctx context.Context) error
	// Callback function
	CallbackPaymentSuccess(id int, ctx context.Context) (string, error)
	CallbackPaymentCancel(id int, ctx context.Context) (string, error)
}
