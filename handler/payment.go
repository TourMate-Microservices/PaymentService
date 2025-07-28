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
// @Param        method query string false "Payment method"
// @Param        customerId query string false "Customer ID"
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
// @Param        page query int false "Page"
// @Param        keyword     query string false "Search keyword"
// @Param        filterProp  query string false "Filter property (e.g. date, price)"
// @Param        order       query string false "Sort order (ASC or DESC)"
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

	id, _ := strconv.Atoi(ctx.Param("id"))

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

// CreatePayment godoc
// @Summary      Create a payment
// @Description  Creates a new payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body request.CreatePaymentRequest true "Create Payment Request"
// @Success 200 {object} entity.Payment "success"
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

	// Create payment and get the created payment back
	payment, err := service.CreatePayment(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    payment,
		Data2:    payment,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// CreatePayosTransaction godoc
// @Summary      Create a PayOS Transaction
// @Description  Initiates a PayOS transaction with the given request body
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        request body request.CreatePayosTransactionRequest true "PayOS Transaction Request"
// @Success      200 {object} response.UrlResponse
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Router       /payment-service/api/v1/payments/create-embedded-payment-link [post]
func CreatePayosTransaction(ctx *gin.Context) {
	var request request.CreatePayosTransactionRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.CreatePayosTransaction(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetPaymentWithService godoc
// @Summary Get payment with service information by ID
// @Description Retrieve a single payment record with service information by its ID
// @Tags payments
// @Produce json
// @Security BearerAuth
// @Param id path int true "Payment ID"
// @Success 200 {object} response.PaymentWithServiceNameResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router /payment-service/api/v1/payments/with-service-name/{id} [get]
func GetPaymentWithService(ctx *gin.Context) {
	service, err := business_logic.GeneratePaymentService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := service.GetPaymentWithService(id, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}
