package repo

import (
	"context"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/entity"
)

type IPaymentRepo interface {
	GetPayments(req request.GetPaymentsRequest, ctx context.Context) (*[]entity.Payment, int, int, error)
	GetPaymentById(id int, ctx context.Context) (*entity.Payment, error)
	CreatePayment(payment entity.Payment, ctx context.Context) (*entity.Payment, error)
	CreatePaymentWithScopeId(payment entity.Payment, ctx context.Context) (int, error)
	UpdatePayment(payment entity.Payment, ctx context.Context) error
}
