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
	CreatePayment(req request.CreatePaymentRequest, ctx context.Context) error
	CreatePayosTransaction(req request.CreatePayosTransactionRequest, ctx context.Context) (response.UrlResponse, error)
	// Callback function
	// CallbackPaymentSuccess(component response.PaymentCallbackComponent, ctx context.Context) (string, error)
	// CallbackPaymentCancel(component response.PaymentCallbackComponent, ctx context.Context) (string, error)
}
