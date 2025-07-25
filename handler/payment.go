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

// GetAllPayments godoc
// @Summary Get all payments
// @Description Retrieve a paginated list of payments
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param        page query int false "Page number"
// @Param        status query string false "Payment status"
// @Param        method query string false "Payment method"
// @Param        user_id query string false "User ID"
// @Success 200 {object} response.PaginationDataResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router /payment-service/api/v1/payments [get]
func GetAllPayments(ctx *gin.Context) {
	var request request.GetPaymentsRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetPayments(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetPaymentById godoc
// @Summary Get payment by ID
// @Description Retrieve a single payment record by its ID
// @Tags payments
// @Produce json
// @Security BearerAuth
// @Param id path int true "Payment ID"
// @Success 200 {object} entity.Payment
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router /payment-service/api/v1/payments/{id} [get]
func GetPaymentById(ctx *gin.Context) {
	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Query("id"))

	res, err := service.GetPaymentById(id, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetPaymentsByUser godoc
// @Summary Get payments by user ID
// @Description Retrieve a list of payments made by a specific customer
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param        id path int true "Customer ID"
// @Param        page_number query int false "Page number"
// @Param        keyword     query string false "Search keyword"
// @Param        filter_prop query string false "Filter property (e.g. date, price)"
// @Param        order       query string false "Sort order (ASC or DESC)"
// @Param        status       query string false "Payment status (e.g. pending, paid)"
// @Param        method       query string false "Payment Method (e.g. VTP, GHTK)"
// @Success 200 {object} response.PaginationDataResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router /payment-service/api/v1/payments/customer/{id} [get]
func GetPaymentsByUser(ctx *gin.Context) {
	var request request.GetPaymentsRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Query("id"))

	request.CustomerId = &id

	res, err := service.GetPayments(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// UpdatePayment godoc
// @Summary Update a payment record
// @Description Update payment details
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.UpdatePaymentRequest true "UpdatePaymentRequest"
// @Success 200 {object} response.MessageApiResponse "Success"
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router /payment-service/api/v1/payments/update [put]
func UpdatePayment(ctx *gin.Context) {
	var request request.UpdatePaymentRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	utils.ProcessResponse(response.ApiResponse{
		ErrMsg:   service.UpdatePayment(request, ctx),
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// CallbackPaymentSuccess godoc
// @Summary      Callback after successful payment
// @Description  Handles redirect or callback after successful payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id path int true "Payment ID"
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /payment-service/api/v1/payments/callback-success/{id} [get]
func CallbackPaymentSuccess(ctx *gin.Context) {
	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Query("id"))

	res, err := service.CallbackPaymentSuccess(id, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.REDIRECT,
	})
}

// CallbackPaymentCancel godoc
// @Summary      Callback after canceled payment
// @Description  Handles redirect or callback after canceled payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id path int true "Payment ID"
// @Failure      400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /payment-service/api/v1/payments/callback-cancel/{id} [get]
func CallbackPaymentCancel(ctx *gin.Context) {
	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Query("id"))

	res, err := service.CallbackPaymentCancel(id, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.REDIRECT,
	})
}

// CreatePayment godoc
// @Summary      Create a payment
// @Description  Creates a new payment and returns redirect info or confirmation
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body request.CreatePaymentRequest true "Create Payment Request"
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/payments/create [post]
func CreatePayment(ctx *gin.Context) {
	var request request.CreatePaymentRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.CreatePayment(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.REDIRECT,
	})
}
