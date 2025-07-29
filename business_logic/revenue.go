package businesslogic

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
	"tourmate/payment-service/constant/noti"
	"tourmate/payment-service/infrastructure/grpc/user"
	"tourmate/payment-service/infrastructure/grpc/user/pb"
	business_logic "tourmate/payment-service/interface/business_logic"
	"tourmate/payment-service/interface/repo"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"
	"tourmate/payment-service/model/entity"
	"tourmate/payment-service/repository"
	"tourmate/payment-service/repository/db"
	db_server "tourmate/payment-service/repository/db_server"
	"tourmate/payment-service/utils"
)

type revenueService struct {
	logger      *log.Logger
	userService business_logic.IUserService
	revenueRepo repo.IRevenueRepo
}

func InitializeRevenueService(db *sql.DB, userService business_logic.IUserService, logger *log.Logger) business_logic.IRevenueService {
	return &revenueService{
		logger:      logger,
		userService: userService,
		revenueRepo: repository.InitializeRevenueRepo(db, logger),
	}
}

func GenerateRevenueService() (business_logic.IRevenueService, error) {
	var logger = utils.GetLogConfig()

	cnn, err := db.ConnectDB(logger, db_server.InitializeMsSQL())

	if err != nil {
		return nil, err
	}

	userService, _ := user.GenerateUserService(logger)

	return InitializeRevenueService(cnn, userService, logger), nil
}

// GetRevenueStats implements businesslogic.IRevenueService.
func (r *revenueService) GetRevenueStats(req request.GetMonthlyRevenueRequest, ctx context.Context) (*response.RevenueStatusResponse, error) {
	// Monthly stats
	revenues, err := r.revenueRepo.GetRevenuesByMonth(req.TourGuideId, req.Year, req.Month, ctx)
	if err != nil {
		return nil, err
	}

	var previousMonth int
	var year int = req.Year
	if req.Month == 1 {
		previousMonth = 12
		year -= 1
	} else {
		previousMonth -= 1
	}
	previousMonthAmount, err := r.revenueRepo.GetRevenueTotalAmountByMonth(req.TourGuideId, year, previousMonth, ctx)
	if err != nil {
		return nil, err
	}

	var totalRevenue, platformFee, netRevenue float64
	var completedPayments, pendingPayments int
	var revenuesResponse []response.RevenueResponse
	var tourguideName string
	if tourguideInfo, _ := r.userService.GetTourGuideById(ctx, &pb.GetTourGuideByIdRequest{
		TourGuideId: int32(req.TourGuideId),
	}); tourguideInfo != nil {
		tourguideName = tourguideInfo.FullName
	}

	for _, rev := range *revenues {
		totalRevenue += rev.TotalAmount
		platformFee += rev.PlatformCommission
		netRevenue += rev.ActualReceived

		if rev.PaymentStatus {
			completedPayments++
		} else {
			pendingPayments++
		}

		revenuesResponse = append(revenuesResponse, response.RevenueResponse{
			RevenueId:          rev.RevenueId,
			PaymentId:          rev.PaymentId,
			InvoiceId:          rev.InvoiceId,
			TourGuideId:        rev.TourGuideId,
			TotalAmount:        rev.TotalAmount,
			ActualReceived:     rev.ActualReceived,
			PlatformCommission: rev.PlatformCommission,
			CreatedAt:          rev.CreatedAt,
			PaymentStatus:      rev.PaymentStatus,
			TourGuideName:      tourguideName,
		})
	}

	var growthPercentage float64
	if previousMonthAmount > 0 {
		growthPercentage = ((totalRevenue - previousMonthAmount) / previousMonthAmount) * 100
	}

	return &response.RevenueStatusResponse{
		TotalRevenue:      totalRevenue,
		PlatformFee:       platformFee,
		NetRevenue:        netRevenue,
		TotalRecords:      len(*revenues),
		CompletedPayments: completedPayments,
		PendingPayments:   pendingPayments,
		RevenueList:       revenuesResponse,
		MonthlyGrowth:     growthPercentage,
	}, nil
}

// GetRevenues implements businesslogic.IRevenueService.
func (r *revenueService) GetRevenues(req request.GetRevenuesRequest, ctx context.Context) (*[]response.RevenueResponse, error) {
	data, err := r.revenueRepo.GetRevenues(req, ctx)

	var tourguideName string
	if tourguideInfo, _ := r.userService.GetTourGuideById(ctx, &pb.GetTourGuideByIdRequest{
		TourGuideId: int32(req.TourGuideId),
	}); tourguideInfo != nil {
		tourguideName = tourguideInfo.FullName
	}

	var res []response.RevenueResponse
	if data != nil {
		for _, item := range *data {
			res = append(res, response.RevenueResponse{
				RevenueId:          item.RevenueId,
				PaymentId:          item.PaymentId,
				InvoiceId:          item.InvoiceId,
				TourGuideId:        item.TourGuideId,
				TotalAmount:        item.TotalAmount,
				ActualReceived:     item.TotalAmount,
				PlatformCommission: item.PlatformCommission,
				CreatedAt:          item.CreatedAt,
				PaymentStatus:      item.PaymentStatus,
				TourGuideName:      tourguideName,
			})
		}
	}

	return &res, err
}

func (r *revenueService) GetMonthlyRevenue(req request.GetMonthlyRevenueRequest, ctx context.Context) (*response.MonthlyRevenueResponse, error) {
	revenues, err := r.revenueRepo.GetRevenuesByMonth(req.TourGuideId, req.Month, req.Year, ctx)
	if err != nil {
		return nil, err
	}

	var previousMonth int
	var year int = req.Year
	if req.Month == 1 {
		previousMonth = 12
		year -= 1
	} else {
		previousMonth -= 1
	}

	previousTotalAmount, err := r.revenueRepo.GetRevenueTotalAmountByMonth(req.TourGuideId, year, previousMonth, ctx)
	if err != nil {
		return nil, err
	}

	var totalRevenue, platformFee, netRevenue float64
	var completedPayments, pendingPayments int

	for _, rev := range *revenues {
		totalRevenue += rev.TotalAmount
		platformFee += rev.PlatformCommission
		netRevenue += rev.ActualReceived

		if rev.PaymentStatus {
			completedPayments++
		} else {
			pendingPayments++
		}
	}

	var growthPercentage float64
	if previousTotalAmount > 0 {
		growthPercentage = ((totalRevenue - previousTotalAmount) / previousTotalAmount) * 100
	}

	return &response.MonthlyRevenueResponse{
		Month:             req.Month,
		Year:              req.Year,
		TotalRevenue:      totalRevenue,
		PlatformFee:       platformFee,
		NetRevenue:        netRevenue,
		TotalRecords:      len(*revenues),
		CompletedPayments: completedPayments,
		PendingPayments:   pendingPayments,
		GrowthPercentage:  growthPercentage,
	}, nil
}

// GetGrowthPercentage implements businesslogic.IRevenueService.
func (r *revenueService) GetGrowthPercentage(req request.GetMonthlyRevenueRequest, ctx context.Context) (response.RevenueGrowthPercentageResponse, error) {
	var previousMonth int
	var year int = req.Year
	if req.Month == 1 {
		previousMonth = 12
		year -= 1
	} else {
		previousMonth -= 1
	}

	previousMonthAmount, err := r.revenueRepo.GetRevenueTotalAmountByMonth(req.TourGuideId, year, previousMonth, ctx)
	if err != nil {
		return response.RevenueGrowthPercentageResponse{}, err
	}

	if previousMonth == 0 {
		return response.RevenueGrowthPercentageResponse{}, nil
	}

	currentMonthAmount, err := r.revenueRepo.GetRevenueTotalAmountByMonth(req.TourGuideId, req.Year, req.Month, ctx)
	if err != nil {
		return response.RevenueGrowthPercentageResponse{}, err
	}

	return response.RevenueGrowthPercentageResponse{
		GrowthPercentage: (currentMonthAmount - previousMonthAmount) / previousMonthAmount * 100,
	}, nil
}

// CreateRevenue implements businesslogic.IRevenueService.
func (r *revenueService) CreateRevenue(req request.CreateRevenueRequest, ctx context.Context) (*response.RevenueResponse, error) {
	var curTime time.Time = time.Now()
	var revenue entity.Revenue = entity.Revenue{
		PaymentId:          req.PaymentId,
		TourGuideId:        req.TourGuideId,
		InvoiceId:          req.InvoiceId,
		TotalAmount:        req.TotalAmount,
		ActualReceived:     req.ActualReceived,
		PlatformCommission: req.PlatformCommission,
		PaymentStatus:      req.PaymentStatus,
		CreatedAt:          curTime,
	}

	id, err := r.revenueRepo.CreateRevenue(revenue, ctx)
	if err != nil {
		return nil, err
	}

	var tourguideName string
	if tourguideInfo, _ := r.userService.GetTourGuideById(ctx, &pb.GetTourGuideByIdRequest{
		TourGuideId: int32(req.TourGuideId),
	}); tourguideInfo != nil {
		tourguideName = tourguideInfo.FullName
	}

	return &response.RevenueResponse{
		RevenueId:          id,
		PaymentId:          req.PaymentId,
		TourGuideId:        req.TourGuideId,
		TourGuideName:      tourguideName,
		InvoiceId:          req.InvoiceId,
		TotalAmount:        req.TotalAmount,
		ActualReceived:     req.ActualReceived,
		PlatformCommission: req.PlatformCommission,
		PaymentStatus:      req.PaymentStatus,
		CreatedAt:          curTime,
	}, nil
}

// GetRevenue implements businesslogic.IRevenueService.
func (r *revenueService) GetRevenue(id int, ctx context.Context) (*entity.Revenue, error) {
	return r.revenueRepo.GetRevenue(id, ctx)
}

// RemoveRevenue implements businesslogic.IRevenueService.
func (r *revenueService) RemoveRevenue(id int, ctx context.Context) error {
	return r.revenueRepo.RemoveRevenue(id, ctx)
}

// UpdateRevenue implements businesslogic.IRevenueService.
func (r *revenueService) UpdateRevenue(req request.UpdateRevenueRequest, ctx context.Context) (*response.RevenueResponse, error) {
	revenue, err := r.revenueRepo.GetRevenue(req.RevenueId, ctx)
	if err != nil {
		return nil, err
	}

	if revenue == nil {
		return nil, errors.New(noti.GENERIC_ERROR_WARN_MSG)
	}

	if req.PaymentId != nil {
		revenue.PaymentId = *req.PaymentId
	}

	if req.TourGuideId != nil {
		revenue.TourGuideId = *req.TourGuideId
	}

	if req.InvoiceId != nil {
		revenue.InvoiceId = *req.InvoiceId
	}

	if req.TotalAmount != nil {
		revenue.TotalAmount = *req.TotalAmount
	}

	if req.ActualReceived != nil {
		revenue.ActualReceived = *req.ActualReceived
	}

	if req.PlatformCommission != nil {
		revenue.PlatformCommission = *req.PlatformCommission
	}

	if req.PaymentStatus != nil {
		revenue.PaymentStatus = *req.PaymentStatus
	}

	if err := r.revenueRepo.UpdateRevenue(*revenue, ctx); err != nil {
		return nil, err
	}

	var tourguideName string
	if tourguideInfo, _ := r.userService.GetTourGuideById(ctx, &pb.GetTourGuideByIdRequest{
		TourGuideId: int32(revenue.TourGuideId),
	}); tourguideInfo != nil {
		tourguideName = tourguideInfo.FullName
	}

	return &response.RevenueResponse{
		RevenueId:          req.RevenueId,
		PaymentId:          revenue.PaymentId,
		TourGuideId:        revenue.TourGuideId,
		TourGuideName:      tourguideName,
		InvoiceId:          revenue.InvoiceId,
		TotalAmount:        revenue.TotalAmount,
		ActualReceived:     revenue.ActualReceived,
		PlatformCommission: revenue.PlatformCommission,
		PaymentStatus:      revenue.PaymentStatus,
		CreatedAt:          revenue.CreatedAt,
	}, nil
}
