package repo

import (
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/entity"

	"golang.org/x/net/context"
)

type IRevenueRepo interface {
	GetRevenues(req request.GetRevenuesRequest, ctx context.Context) (*[]entity.Revenue, error)
	GetRevenuesByMonth(tourGuideId, year, month int, ctx context.Context) (*[]entity.Revenue, error)
	GetRevenueTotalAmountByMonth(tourGuideId, year, month int, ctx context.Context) (float64, error)
	GetCountTotalRevenue(req request.GetRevenuesRequest, ctx context.Context) (int, error)
	GetRevenue(id int, ctx context.Context) (*entity.Revenue, error)
	CreateRevenue(revenue entity.Revenue, ctx context.Context) (int, error)
	UpdateRevenue(revenue entity.Revenue, ctx context.Context) error
	RemoveRevenue(id int, ctx context.Context) error
}
