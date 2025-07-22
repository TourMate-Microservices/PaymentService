package api

import (
	"tourmate/payment-service/handler"
	"tourmate/payment-service/utils/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeFeedbackHandlerRoute(server *gin.Engine, port, service string) {
	//Context path
	var contextPath string = service + "api/v1/feedbacks"

	// Define Feedback endpoints with basic required
	var authGroup = server.Group(contextPath, middleware.Authorize)
	authGroup.GET("", handler.GetFeedbacks)
	authGroup.GET("/user/:id", handler.GetFeedbacksByUser)
	authGroup.GET("/:id", handler.GetFeedbackById)
	authGroup.POST("", handler.CreateFeedback)
	authGroup.PUT("", handler.UpdateFeedback)
	authGroup.DELETE("", handler.RemoveFeedback)
}
