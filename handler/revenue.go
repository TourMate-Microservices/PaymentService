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

// GetRevenues godoc
// @Summary      Get all revenue entries
// @Description  Retrieves revenues with optional filters
// @Tags         revenues
// @Accept       json
// @Produce      json
// @Param        query query request.GetRevenuesRequest true "Revenue Query Params"
// @Success      200 {array} response.RevenueResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/revenues [get]
// @Security     BearerAuth
func GetRevenues(ctx *gin.Context) {
	var request request.GetRevenuesRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateRevenueService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.GetRevenues(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetMonthlyRevenue godoc
// @Summary      Get revenue by month
// @Description  Retrieves revenue grouped by month
// @Tags         revenues
// @Accept       json
// @Produce      json
// @Param        id path int true "Tour Guide ID"
// @Param        query query request.GetMonthlyRevenueRequest true "Monthly Revenue Query"
// @Success      200 {object} response.MonthlyRevenueResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/revenues/monthly/{id} [get]
// @Security     BearerAuth
func GetMonthlyRevenue(ctx *gin.Context) {
	var request request.GetMonthlyRevenueRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateRevenueService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	request.TourGuideId = id

	res, err := service.GetMonthlyRevenue(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetGrowthPercentage godoc
// @Summary      Get growth percentage
// @Description  Calculates revenue growth over time
// @Tags         revenues
// @Accept       json
// @Produce      json
// @Param        id path int true "Tour Guide ID"
// @Param        query query request.GetMonthlyRevenueRequest true "Growth Percentage Query Params"
// @Success      200 {object} response.RevenueGrowthPercentageResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/revenues/growth/{id} [get]
// @Security     BearerAuth
func GetGrowthPercentage(ctx *gin.Context) {
	var request request.GetMonthlyRevenueRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateRevenueService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	request.TourGuideId = id

	res, err := service.GetGrowthPercentage(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetRevenue godoc
// @Summary      Get a single revenue record
// @Description  Retrieves revenue details by ID
// @Tags         revenues
// @Accept       json
// @Produce      json
// @Param        id path int true "Revenue ID"
// @Success      200 {object} entity.Revenue
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/revenues/{id} [get]
// @Security     BearerAuth
func GetRevenue(ctx *gin.Context) {
	service, err := business_logic.GenerateRevenueService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := service.GetRevenue(id, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// CreateRevenue godoc
// @Summary      Create a revenue record
// @Description  Adds a new revenue entry
// @Tags         revenues
// @Accept       json
// @Produce      json
// @Param        request body request.CreateRevenueRequest true "Create Revenue Payload"
// @Success      200 {object} response.RevenueResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/revenues [post]
// @Security     BearerAuth
func CreateRevenue(ctx *gin.Context) {
	var request request.CreateRevenueRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateRevenueService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	res, err := service.CreateRevenue(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.CREATE_ACTION,
	})
}

// UpdateRevenue godoc
// @Summary      Update a revenue record
// @Description  Updates an existing revenue entry
// @Tags         revenues
// @Accept       json
// @Produce      json
// @Param        id path int true "Revenue ID"
// @Param        request body request.UpdateRevenueRequest true "Update Revenue Payload"
// @Success      200 {object} response.RevenueResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/revenues/{id} [put]
// @Security     BearerAuth
func UpdateRevenue(ctx *gin.Context) {
	var request request.UpdateRevenueRequest
	if ctx.ShouldBindJSON(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateRevenueService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	request.RevenueId = id

	res, err := service.UpdateRevenue(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// RemoveRevenue godoc
// @Summary      Delete a revenue record
// @Description  Removes a revenue entry by ID
// @Tags         revenues
// @Accept       json
// @Produce      json
// @Param        id path int true "Revenue ID"
// @Success 200 {object} response.MessageApiResponse "Success"
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/revenues{id} [delete]
// @Security     BearerAuth
func RemoveRevenue(ctx *gin.Context) {
	service, err := business_logic.GenerateRevenueService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	utils.ProcessResponse(response.ApiResponse{
		ErrMsg:   service.RemoveRevenue(id, ctx),
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}

// GetMonthlyRevenue godoc
// @Summary      Get revenue stats
// @Description  Retrieves revenue stats
// @Tags         revenues
// @Accept       json
// @Produce      json
// @Param        id path int true "Revenue ID"
// @Param        query query request.GetMonthlyRevenueRequest true "Revenue Stats Query"
// @Success      200 {object} response.RevenueStatusResponse
// @Failure 401 {object} response.MessageApiResponse "You have no rights to access this action."
// @Failure 400 {object} response.MessageApiResponse "Invalid data. Please try again."
// @Failure 500 {object} response.MessageApiResponse "There is something wrong in the system during the process. Please try again later."
// @Router       /payment-service/api/v1/revenues/stats/{id} [get]
// @Security     BearerAuth
func GetRevenueStats(ctx *gin.Context) {
	var request request.GetMonthlyRevenueRequest
	if ctx.ShouldBindQuery(&request) != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, nil))
		return
	}

	service, err := business_logic.GenerateRevenueService()
	if err != nil {
		utils.ProcessResponse(utils.GenerateInvalidRequestAndSystemProblemModel(ctx, err))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	request.TourGuideId = id

	res, err := service.GetRevenueStats(request, ctx)

	utils.ProcessResponse(response.ApiResponse{
		Data1:    res,
		Data2:    res,
		ErrMsg:   err,
		Context:  ctx,
		PostType: action_type.NON_POST,
	})
}
