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
	user_pb "tourmate/payment-service/infrastructure/grpc/user/pb"

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
func (f *feedbackService) GetTourGuideFeedbacks(req request.GetTourGuideFeedbacksRequest, ctx context.Context) (response.PaginationDataResponse, error) {
	if req.TourGuideId < 1 {
		req.TourGuideId = 1
	}

	f.logger.Printf("GetTourGuideFeedbacks: Starting with TourGuideId=%d, PageIndex=%d, PageSize=%d", req.TourGuideId, req.PageIndex, req.PageSize)

	feedbacks, pages, totalRecords, err := f.feedbackRepo.GetFeedbacks(request.GetFeedbacksRequest{
		Request: request.SearchPaginationRequest{
			Page:       req.PageIndex,
			FilterProp: assignFilterProperty(""),
			Order:      utils.AssignOrder(""),
		},
		PageSize:    req.PageSize,
		TourGuideId: &req.TourGuideId,
	}, ctx)

	if err != nil {
		f.logger.Printf("GetTourGuideFeedbacks: Database error: %v", err)
		return response.PaginationDataResponse{}, err
	}

	f.logger.Printf("GetTourGuideFeedbacks: Found %d feedbacks", len(*feedbacks))

	var data []response.FeedbackResponse
	for i, feedback := range *feedbacks {
		f.logger.Printf("GetTourGuideFeedbacks: Processing feedback %d with CustomerId=%d, ServiceId=%d", i+1, feedback.CustomerId, feedback.ServiceId)

		// Initialize response with basic feedback data
		feedbackResponse := response.FeedbackResponse{
			FeedbackId:  feedback.FeedbackId,
			CustomerId:  feedback.CustomerId,
			FullName:    "Unknown Customer", // Default value
			Image:       "",                 // Default value
			Rating:      feedback.Rating,
			Content:     feedback.Content,
			CreatedDate: feedback.CreatedDate,
			InvoiceId:   feedback.InvoiceId,
			ServiceId:   feedback.ServiceId,
			ServiceName: "Unknown Service", // Default value
		}

		// Try to get customer information via gRPC
		customerInfo, err := f.userService.GetCustomerById(ctx, &user_pb.GetCustomerByIdRequest{
			CustomerId: int32(feedback.CustomerId),
		})

		if err != nil {
			f.logger.Printf("GetTourGuideFeedbacks: UserService gRPC call failed for CustomerId=%d: %v", feedback.CustomerId, err)
		} else if customerInfo != nil {
			f.logger.Printf("GetTourGuideFeedbacks: Got customer info for CustomerId=%d: %s", feedback.CustomerId, customerInfo.FullName)
			feedbackResponse.FullName = customerInfo.FullName
			feedbackResponse.Image = customerInfo.Image
		}

		// Try to get tour information via gRPC
		tourInfo, err := f.tourService.GetTourById(ctx, &tour_pb.TourServiceIdRequest{
			ServiceId: int32(feedback.ServiceId),
		})

		if err != nil {
			f.logger.Printf("GetTourGuideFeedbacks: TourService gRPC call failed for ServiceId=%d: %v", feedback.ServiceId, err)
		} else if tourInfo != nil {
			f.logger.Printf("GetTourGuideFeedbacks: Got tour info for ServiceId=%d: %s", feedback.ServiceId, tourInfo.ServiceName)
			feedbackResponse.ServiceName = tourInfo.ServiceName
		}

		data = append(data, feedbackResponse)
	}

	f.logger.Printf("GetTourGuideFeedbacks: Successfully processed %d feedbacks", len(data))

	return response.PaginationDataResponse{
		Data:        data,
		TotalCount:  totalRecords,
		Page:        req.PageIndex,
		PerPage:     req.PageSize,
		TotalPages:  pages,
		HasNext:     req.PageIndex < pages,
		HasPrevious: req.PageIndex > 1,
	}, nil
}

// CreateFeedback implements businesslogic.IFeedbackService.
func (f *feedbackService) CreateFeedback(req request.CreateFeedbackRequest, ctx context.Context) error {
	// Verify user data (implement later)
	user, err := f.userService.GetCustomerById(ctx, &user_pb.GetCustomerByIdRequest{
		CustomerId: int32(req.CustomerId),
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

	req.PageSize = req.Request.Page

	data, pages, totalRecords, err := f.feedbackRepo.GetFeedbacks(req, ctx)

	return response.PaginationDataResponse{
		Data:        data,
		TotalCount:  totalRecords,
		Page:        req.Request.Page,
		PerPage:     10,
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
