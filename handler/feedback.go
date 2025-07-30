package handler

import (
	"os"
	"strconv"
	business_logic "tourmate/payment-service/business_logic"
	action_type "tourmate/payment-service/constant/action_type"
	"tourmate/payment-service/constant/env"
	"tourmate/payment-service/infrastructure/grpc/feedback/pb"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"
	"tourmate/payment-service/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetFeedbacks godoc
// @Summary      Get all feedbacks
// @Description  Retrieve a paginated list of feedbacks with optional filters
// @Tags         feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Page"
// @Param        keyword     query string false "Search keyword"
// @Param        filterProp  query string false "Filter property"
// @Param        order       query string false "Sort order (ASC or DESC)"
// @Param        rating      query int false "Rating"
// @Param        customerId  query int false "The owner ID of this feedback"
// @Param        tourGuideId query int false "Tour guide ID"
// @Param        invoiceId   query int false "Invoice ID"
// @Success      200 {object} response.PaginationDataResponse
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /payment-service/api/v1/feedbacks [get]
func GetFeedbacks(ctx *gin.Context) {
	var request request.GetFeedbacksRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetFeedbacks(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetFeedbacksByUser godoc
// @Summary      Get feedbacks by user
// @Description  Retrieve feedbacks filtered by customer ID
// @Tags         feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id          path int  true  "Customer ID"
// @Param        page query int false "Page"
// @Param        keyword     query string false "Search keyword"
// @Param        filterProp query string false "Filter property"
// @Param        order       query string false "Sort order (ASC or DESC)"
// @Param        rating      query int false "Rating"
// @Param        tourGuideId query int false "Tour guide ID"
// @Param        invoiceId   query int false "Invoice ID"
// @Success      200 {object} response.PaginationDataResponse
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /payment-service/api/v1/feedbacks/user/{id} [get]
func GetFeedbacksByUser(ctx *gin.Context) {
	var request request.GetFeedbacksRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	customerId, _ := strconv.Atoi(ctx.Param("id"))

	request.CustomerId = &customerId

	res, err := service.GetFeedbacks(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetFeedbackById godoc
// @Summary      Get a single feedback by ID
// @Tags         feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path int true "Feedback ID"
// @Success      200 {object} entity.Feedback
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /payment-service/api/v1/feedbacks/{id} [get]
func GetFeedbackById(ctx *gin.Context) {
	service, err := business_logic.GenerateFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := service.GetFeedbackById(id, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// CreateFeedback godoc
// @Summary      Create a new feedback
// @Tags         feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body request.CreateFeedbackRequest true "Create Feedback Request"
// @Success 200 {object} response.MessageApiResponse "success"
// @Success 201 {object} entity.Feedback
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/feedbacks [post]
func CreateFeedback(ctx *gin.Context) {
	var request request.CreateFeedbackRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.CreateFeedback(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.CREATE_ACTION,
	})
}

// UpdateFeedback godoc
// @Summary      Update an existing feedback
// @Tags         feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body request.UpdateFeedbackRequest true "Update Feedback Request"
// @Success 200 {object} response.MessageApiResponse "success"
// @Success 201 {object} entity.Feedback
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/feedbacks [put]
func UpdateFeedback(ctx *gin.Context) {
	var request request.UpdateFeedbackRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.UpdateFeedback(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// RemoveFeedback godoc
// @Summary      Remove a feedback
// @Tags         feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body request.RemoveFeedbackRequest true "Remove Feedback Request"
// @Success 200 {object} response.MessageApiResponse "success"
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/feedbacks [delete]
func RemoveFeedback(ctx *gin.Context) {
	var request request.RemoveFeedbackRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	utils.ProcessResponse(response.ApiResponse{
		ErrMsg:   service.RemoveFeedback(request, ctx),
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// TestGrpcFeedback godoc
// @Summary      Test gRPC call to get payment service rating
// @Description  Calls the gRPC feedback service to retrieve average rating and total count for a tour service
// @Tags         test-grpc
// @Accept       json
// @Produce      json
// @Param        id path int true "Tour service ID"
// @Success      200 {object} pb.TourServiceRatingResponse
// @Failure      401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure      500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/feedbacks/test-grpc/{id} [get]
func TestGrpcFeedback(ctx *gin.Context) {
	cnn, err := grpc.Dial("localhost:"+os.Getenv(env.PAYMENT_SERVICE_GRPC_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Method này ko đc thì undo commeent ở trên và thử lại
	// cnn, err := grpc.NewClient("localhost:"+os.Getenv(env.PAYMENT_SERVICE_GRPC_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}
	defer cnn.Close()

	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := pb.NewPaymentServiceClient(cnn).GetTourServiceRating(ctx, &pb.GetTourServiceRatingRequest{
		ServiceId: int32(id),
	})

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetTourGuideFeedbacks godoc
// @Summary      Get feedbacks of a tour guide
// @Description  Retrieves a paginated list of feedbacks to a specific tour guide
// @Tags         feedbacks
// @Accept       json
// @Produce      json
// @Param        id path int true "Tour Guide ID"
// @Param        pageIndex  query int  false  "Page Index for pagination (starts from 1)"
// @Param        pageSize   query     int  true  "Number of items per page"
// @Success      200 {object} response.PaginationDataResponse
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /payment-service/api/v1/feedbacks/tourGuide/{id} [get]
func GetTourGuideFeedbacks(ctx *gin.Context) {
	var request request.GetTourGuideFeedbacksRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	tourGuideId, _ := strconv.Atoi(ctx.Param("id"))
	request.TourGuideId = tourGuideId

	// Set default values if not provided
	if request.PageSize <= 0 {
		request.PageSize = 10
	}
	if request.PageIndex <= 0 {
		request.PageIndex = 1
	}

	// Initialize business logic service
	service, err := business_logic.GenerateFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	// Call GetTourGuideFeedbacks (with graceful gRPC fallback handling)
	res, err := service.GetTourGuideFeedbacks(request, ctx)
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   nil,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}
