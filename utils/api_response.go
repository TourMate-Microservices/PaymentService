package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	action_type "tourmate/payment-service/constant/action_type"
	"tourmate/payment-service/constant/noti"
	"tourmate/payment-service/model/dto/response"

	"github.com/gin-gonic/gin"
)

func ProcessResponse(data response.ApiResponse) {
	if data.ErrMsg != nil {
		processFailResponse(data.ErrMsg, data.Context)
		return
	}

	if data.PostType != action_type.NON_POST {
		processSuccessPostReponse(data.Data2, data.PostType, data.Context)
		return
	}

	processSuccessResponse(data.Data1, data.Context)
}

func GenerateInvalidRequestAndSystemProblemModel(ctx *gin.Context, err error) response.ApiResponse {
	var errMsg error = err
	if errMsg == nil {
		errMsg = errors.New(noti.GENERIC_ERROR_WARN_MSG)
	}

	return response.ApiResponse{
		ErrMsg:   errMsg,
		Context:  ctx,
		PostType: action_type.NON_POST,
	}
}

func GetUnAuthBodyResponse(ctx *gin.Context) response.ApiResponse {
	return response.ApiResponse{
		ErrMsg:  errors.New(noti.GENERIC_RIGHT_ACCESS_WARN_MSG),
		Context: ctx,
	}
}

func processFailResponse(err error, ctx *gin.Context) {
	var errCode int

	switch err.Error() {
	case noti.INTERNALL_ERR_MSG:
		errCode = http.StatusInternalServerError
	case noti.GENERIC_RIGHT_ACCESS_WARN_MSG:
		errCode = http.StatusForbidden
	default:
		errCode = http.StatusBadRequest
	}

	if isErrorTypeOfUndefined(err) {
		errCode = http.StatusNotFound
	}

	ctx.IndentedJSON(errCode, response.MessageApiResponse{
		Message: err.Error(),
	})
}

func processSuccessPostReponse(res interface{}, postType string, ctx *gin.Context) {
	switch postType {
	case action_type.REDIRECT:
		processRedirectResponse(fmt.Sprint(res), ctx)
	case action_type.INFORM:
		processInformResponse(res, ctx)
	default:
		ctx.IndentedJSON(http.StatusOK, response.MessageApiResponse{
			Message: "success",
		})
	}
}

func processRedirectResponse(redirectUrl string, ctx *gin.Context) {
	ctx.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

func processInformResponse(message interface{}, ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, response.MessageApiResponse{
		Message: fmt.Sprint(message),
	})
}

func processSuccessResponse(data interface{}, ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, data)
}

func isErrorTypeOfUndefined(err error) bool {
	return strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "undefined")
}
