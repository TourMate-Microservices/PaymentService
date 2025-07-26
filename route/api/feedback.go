package api

import (
	"tourmate/payment-service/handler"

	"github.com/gin-gonic/gin"
)

func InitializeFeedbackHandlerRoute(server *gin.Engine, port, service string) {
	//Context path
	var contextPath string = service + "/api/v1/feedbacks"

	// Define Feedback endpoints with basic required
	var authGroup = server.Group(contextPath)
	authGroup.GET("", handler.GetFeedbacks)
	authGroup.GET("/user/:id", handler.GetFeedbacksByUser)
	authGroup.GET("/:id", handler.GetFeedbackById)
	authGroup.POST("", handler.CreateFeedback)
	authGroup.PUT("", handler.UpdateFeedback)
	authGroup.DELETE("", handler.RemoveFeedback)
	authGroup.GET("/test-grpc/:id", handler.TestGrpcFeedback)
}
