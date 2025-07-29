package api

import (
	"tourmate/payment-service/handler"

	"github.com/gin-gonic/gin"
)

func InitializePlatformFeedbackHandlerRoute(server *gin.Engine, service string) {
	//Context path
	var contextPath string = "api/v1/platform-feedbacks"

	// Define Feedback endpoints with basic required
	var authGroup = server.Group(contextPath)
	authGroup.GET("", handler.GetPlatformFeedbacks)
	authGroup.GET("/customer/:id", handler.GetPlatformFeedbacksByCustomer)
	authGroup.GET("/:id", handler.GetPlatformFeedbackById)
	authGroup.POST("", handler.CreatePlatformFeedback)
	authGroup.PUT("", handler.UpdatePlatformFeedback)
}
