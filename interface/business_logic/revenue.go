package businesslogic

import (
	"context"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"
	"tourmate/payment-service/model/entity"
)

type IRevenueService interface {
	GetRevenues(req request.GetRevenuesRequest, ctx context.Context) (*[]response.RevenueResponse, error)
	GetMonthlyRevenue(req request.GetMonthlyRevenueRequest, ctx context.Context) (*response.MonthlyRevenueResponse, error)
	GetRevenueStats(req request.GetMonthlyRevenueRequest, ctx context.Context) (*response.RevenueStatusResponse, error)
	GetGrowthPercentage(req request.GetMonthlyRevenueRequest, ctx context.Context) (response.RevenueGrowthPercentageResponse, error)
	GetRevenue(id int, ctx context.Context) (*entity.Revenue, error)
	CreateRevenue(req request.CreateRevenueRequest, ctx context.Context) (*response.RevenueResponse, error)
	UpdateRevenue(req request.UpdateRevenueRequest, ctx context.Context) (*response.RevenueResponse, error)
	RemoveRevenue(id int, ctx context.Context) error
}
