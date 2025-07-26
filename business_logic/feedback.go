package businesslogic

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
	filter_property "tourmate/payment-service/constant/filter_property"
	"tourmate/payment-service/constant/noti"
	"tourmate/payment-service/infrastructure/grpc/tour"
	tour_pb "tourmate/payment-service/infrastructure/grpc/tour/pb"
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

type feedbackService struct {
	logger       *log.Logger
	userService  business_logic.IUserService
	tourService  business_logic.ITourService
	feedbackRepo repo.IFeedbackRepo
}

func InitializeFeedbackService(db *sql.DB, userService business_logic.IUserService, tourService business_logic.ITourService, logger *log.Logger) business_logic.IFeedbackService {
	return &feedbackService{
		logger:       logger,
		userService:  userService,
		tourService:  tourService,
		feedbackRepo: repository.InitializeFeedbackRepo(db, logger),
	}
}

func GenerateFeedbackService() (business_logic.IFeedbackService, error) {
	var logger = utils.GetLogConfig()

	dbCnn, err := db.ConnectDB(logger, db_server.InitializeMsSQL())

	if err != nil {
		return nil, err
	}

	userService, _ := user.GenerateUserService(logger)
	tourService, _ := tour.GenerateTourService(logger)

	return InitializeFeedbackService(dbCnn, userService, tourService, logger), nil
}

// GetFeedbacksUiResponse implements businesslogic.IFeedbackService.
func (f *feedbackService) GetTourGuideFeedbacks(tourGuideId, page int, ctx context.Context) (response.PaginationDataResponse, error) {
	if page < 1 {
		page = 1
	}

	if tourGuideId < 1 {
		tourGuideId = 1
	}

	feedbacks, pages, totalRecords, err := f.feedbackRepo.GetFeedbacks(request.GetFeedbacksRequest{
		Request: request.SearchPaginationRequest{
			Page: page,
		},
		TourGuideId: &tourGuideId,
	}, ctx)

	if err != nil {
		return response.PaginationDataResponse{}, err
	}

	var data []response.FeedbackResponse
	for _, feedback := range *feedbacks {
		customerInfo, err := f.userService.GetUser(ctx, &pb.GetUserRequest{
			Id: int32(feedback.CustomerId),
		})

		if err != nil {
			return response.PaginationDataResponse{}, err
		}

		tourInfo, err := f.tourService.GetTour(ctx, &tour_pb.GetTourRequest{
			Id: int32(feedback.ServiceId),
		})

		if err != nil {
			return response.PaginationDataResponse{}, err
		}

		data = append(data, response.FeedbackResponse{
			FeedbackId:  feedback.FeedbackId,
			CustomerId:  feedback.CustomerId,
			FullName:    customerInfo.Fullname,
			Image:       customerInfo.Image,
			Rating:      feedback.Rating,
			Content:     feedback.Content,
			CreatedDate: feedback.CreatedDate,
			ServiceId:   feedback.ServiceId,
			ServiceName: tourInfo.ServiceName,
		})
	}

	return response.PaginationDataResponse{
		Data:        data,
		TotalCount:  totalRecords,
		Page:        page,
		PerPage:     entity.Feedback{}.GetFeedbackLimitRecords(),
		TotalPages:  pages,
		HasNext:     page < pages,
		HasPrevious: page > 1,
	}, err
}

// CreateFeedback implements businesslogic.IFeedbackService.
func (f *feedbackService) CreateFeedback(req request.CreateFeedbackRequest, ctx context.Context) error {
	// Verify user data (implement later)
	user, err := f.userService.GetUser(ctx, &pb.GetUserRequest{
		Id: int32(req.CustomerId),
	})

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New(noti.GENERIC_ERROR_WARN_MSG)
	}

	// Insert to database
	var curTime time.Time = time.Now()
	return f.feedbackRepo.CreateFeedback(entity.Feedback{
		CustomerId:  req.CustomerId,
		ServiceId:   req.ServiceId,
		TourGuideId: req.TourGuideId,
		InvoiceId:   req.InvoiceId,
		Content:     req.Content,
		Rating:      req.Rating,
		IsDeleted:   false,
		CreatedDate: curTime,
		UpdatedAt:   curTime,
	}, ctx)
}

// GetFeedbackById implements businesslogic.IFeedbackService.
func (f *feedbackService) GetFeedbackById(id int, ctx context.Context) (*entity.Feedback, error) {
	return f.feedbackRepo.GetFeedbackById(id, ctx)
}

// GetFeedbacks implements businesslogic.IFeedbackService.
func (f *feedbackService) GetFeedbacks(req request.GetFeedbacksRequest, ctx context.Context) (response.PaginationDataResponse, error) {
	if req.Request.Page < 1 {
		req.Request.Page = 1
	}

	req.Request.FilterProp = assignFilterProperty(req.Request.FilterProp)
	req.Request.Order = utils.AssignOrder(req.Request.Order)

	data, pages, totalRecords, err := f.feedbackRepo.GetFeedbacks(req, ctx)

	return response.PaginationDataResponse{
		Data:        data,
		TotalCount:  totalRecords,
		Page:        req.Request.Page,
		PerPage:     entity.Feedback{}.GetFeedbackLimitRecords(),
		TotalPages:  pages,
		HasNext:     req.Request.Page < pages,
		HasPrevious: req.Request.Page > 1,
	}, err
}

// RemoveFeedback implements businesslogic.IFeedbackService.
func (f *feedbackService) RemoveFeedback(req request.RemoveFeedbackRequest, ctx context.Context) error {
	feedback, err := f.feedbackRepo.GetFeedbackById(req.FeedbackId, ctx)
	if err != nil {
		return err
	}

	// Verify user data (implement later)

	feedback.IsDeleted = true
	feedback.UpdatedAt = time.Now()

	return f.feedbackRepo.UpdateFeedback(*feedback, ctx)
}

// UpdateFeedback implements businesslogic.IFeedbackService.
func (f *feedbackService) UpdateFeedback(req request.UpdateFeedbackRequest, ctx context.Context) error {
	feedback, err := f.feedbackRepo.GetFeedbackById(req.Request.FeedbackId, ctx)
	if err != nil {
		return err
	}

	// Verify user data (implement later)
	if req.Content != "" {
		feedback.Content = req.Content
	}

	if req.Rating != nil && *req.Rating >= 1 {
		feedback.Rating = *req.Rating
	}

	feedback.UpdatedAt = time.Now()

	return f.feedbackRepo.UpdateFeedback(*feedback, ctx)
}

func assignFilterProperty(filterProp string) string {
	var res string

	switch filterProp {
	case filter_property.DATE_FILTER:
		res = "createdDate"
	case filter_property.ACTION_DATE_FILTER:
		res = "date"
	case filter_property.PRICE_FILTER:
		res = "price"
	case filter_property.RATE_FILTER:
		res = "rate"
	case filter_property.AMOUNT_FILTER:
		res = "amount"
	default:
		res = "createdDate"
	}

	return res
}
