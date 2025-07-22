package repo

import (
	"context"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/entity"
)

type IPaymentRepo interface {
	GetPayments(req request.GetPaymentsRequest, ctx context.Context) (*[]entity.Payment, int, error)
	GetPaymentById(id int, ctx context.Context) (*entity.Payment, error)
	CreatePayment(payment entity.Payment, ctx context.Context) error
	UpdatePayment(payment entity.Payment, ctx context.Context) error
}
