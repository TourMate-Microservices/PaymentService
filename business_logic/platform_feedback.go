package businesslogic

import (
	"context"
	"database/sql"
	"log"
	"time"
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

type platformFeedbackService struct {
	logger               *log.Logger
	platformFeedbackRepo repo.IPlatformFeedbackRepo
}

func InitializePlatformFeedbackService(db *sql.DB, logger *log.Logger) business_logic.IPlatformFeedbackService {
	return &platformFeedbackService{
		logger:               logger,
		platformFeedbackRepo: repository.InitializePlatformFeedbackRepo(db, logger),
	}
}

func GeneratePlatformFeedbackService() (business_logic.IPlatformFeedbackService, error) {
	var logger = utils.GetLogConfig()

	dbCnn, err := db.ConnectDB(logger, db_server.InitializeMsSQL())

	if err != nil {
		return nil, err
	}

	return InitializePlatformFeedbackService(dbCnn, logger), nil
}

// CreatePlatformFeedback implements businesslogic.IPlatformFeedbackService.
func (p *platformFeedbackService) CreatePlatformFeedback(req request.CreatePlatformFeedbackRequest, ctx context.Context) error {
	// Verify user data (implement later)

	// Insert to database
	return p.platformFeedbackRepo.CreatePlatformFeedback(entity.PlatformFeedback{
		AccountId: req.AccountId,
		PaymentId: req.PaymentId,
		Content:   req.Content,
		Rating:    req.Rating,
		CreatedAt: time.Now(),
	}, ctx)
}

// GetPlatformFeedbackById implements businesslogic.IPlatformFeedbackService.
func (p *platformFeedbackService) GetPlatformFeedbackById(id int, ctx context.Context) (*entity.PlatformFeedback, error) {
	return p.platformFeedbackRepo.GetPlatformFeedbackById(id, ctx)
}

// GetPlatformFeedbacks implements businesslogic.IPlatformFeedbackService.
func (p *platformFeedbackService) GetPlatformFeedbacks(req request.GetPlatformFeedbacksRequest, ctx context.Context) (response.PaginationDataResponse, error) {
	if req.Request.Page < 1 {
		req.Request.Page = 1
	}

	req.Request.FilterProp = utils.AssignFilterProperty(req.Request.FilterProp)
	req.Request.Order = utils.AssignOrder(req.Request.Order)

	data, pages, totalRecords, err := p.platformFeedbackRepo.GetPlatformFeedbacks(req, ctx)

	return response.PaginationDataResponse{
		Data:        data,
		TotalCount:  totalRecords,
		Page:        req.Request.Page,
		PerPage:     entity.PlatformFeedback{}.GetPlatformFeedbackLimitRecords(),
		TotalPages:  pages,
		HasNext:     req.Request.Page < pages,
		HasPrevious: req.Request.Page > 1,
	}, err
}

// UpdatePlatformFeedback implements businesslogic.IPlatformFeedbackService.
func (p *platformFeedbackService) UpdatePlatformFeedback(req request.UpdatePlatformFeedbackRequest, ctx context.Context) error {
	platformFeedback, err := p.platformFeedbackRepo.GetPlatformFeedbackById(req.FeedbackId, ctx)
	if err != nil {
		return err
	}

	// Verify user data (implement later)

	if req.Content != "" {
		platformFeedback.Content = req.Content
	}

	if req.Rating != nil && *req.Rating >= 1 {
		platformFeedback.Rating = *req.Rating
	}

	return p.platformFeedbackRepo.UpdatePlatformFeedback(*platformFeedback, ctx)
}
