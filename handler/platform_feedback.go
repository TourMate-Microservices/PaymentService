package handler

import (
	"strconv"
	business_logic "tourmate/payment-service/business_logic"
	action_type "tourmate/payment-service/constant/action_type"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"
	"tourmate/payment-service/utils"

	"github.com/gin-gonic/gin"
)

// GetFeedbacks godoc
// @Summary      Get all platform feedbacks
// @Description  Retrieve a paginated list of platform feedbacks with optional filters
// @Tags         platform-feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        keyword     query string false "Search keyword"
// @Param        filterProp query string false "Filter property"
// @Param        order       query string false "Sort order (ASC or DESC)"
// @Param        customerId  query int false "Customer ID"
// @Param        rating      query int false "Rating"
// @Param        pageIndex  query int false "Page index"
// @Param        pageSize   query int false "Page size"
// @Success      200 {object} response.PaginationDataResponse
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /api/v1/platform-feedbacks [get]
func GetPlatformFeedbacks(ctx *gin.Context) {
	var request request.GetPlatformFeedbacksRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePlatformFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetPlatformFeedbacks(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetFeedbacksByUser godoc
// @Summary      Get platform feedbacks by user
// @Description  Retrieve platform feedbacks filtered by customer ID
// @Tags         platform-feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id          path int  true  "Customer ID"
// @Param        keyword     query string false "Search keyword"
// @Param        filterProp query string false "Filter property"
// @Param        order       query string false "Sort order (ASC or DESC)"
// @Param        rating      query int false "Rating"
// @Param        pageIndex  query int false "Page index"
// @Param        pageSize   query int false "Page size"
// @Success      200 {object} response.PaginationDataResponse
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /api/v1/platform-feedbacks/customer/{id} [get]
func GetPlatformFeedbacksByCustomer(ctx *gin.Context) {
	var request request.GetPlatformFeedbacksRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePlatformFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	customerId, _ := strconv.Atoi(ctx.Param("id"))

	request.CustomerId = &customerId

	res, err := service.GetPlatformFeedbacks(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetFeedbackById godoc
// @Summary      Get a single platform feedback by ID
// @Tags         platform-feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path int true "Platform Feedback ID"
// @Success      200 {object} entity.PlatformFeedback
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /api/v1/platform-feedbacks/{id} [get]
func GetPlatformFeedbackById(ctx *gin.Context) {
	service, err := business_logic.GeneratePlatformFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := service.GetPlatformFeedbackById(id, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// CreatePlatformFeedback godoc
// @Summary      Create a new feedback
// @Tags         platform-feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body request.CreatePlatformFeedbackRequest true "Create Platform Feedback Request"
// @Success 200 {object} response.MessageApiResponse "success"
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /api/v1/platform-feedbacks [post]
func CreatePlatformFeedback(ctx *gin.Context) {
	var request request.CreatePlatformFeedbackRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePlatformFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	utils.ProcessResponse(response.ApiResponse{
		ErrMsg:   service.CreatePlatformFeedback(request, ctx),
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// UpdatePlatformFeedback godoc
// @Summary      Update an existing feedback
// @Tags         platform-feedbacks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body request.UpdatePlatformFeedbackRequest true "Update Feedback Request"
// @Success 200 {object} response.MessageApiResponse "success"
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /api/v1/platform-feedbacks [put]
func UpdatePlatformFeedback(ctx *gin.Context) {
	var request request.UpdatePlatformFeedbackRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePlatformFeedbackService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	utils.ProcessResponse(response.ApiResponse{
		ErrMsg:   service.UpdatePlatformFeedback(request, ctx),
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}
