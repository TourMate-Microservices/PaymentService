package api

import (
	"os"
	"tourmate/payment-service/handler"

	"github.com/gin-gonic/gin"
)

func InitializeFeedbackHandlerRoute(server *gin.Engine, service string) {
	//Context path
	var contextPath string
	if os.Getenv("DOCKER_COMPOSE") == "true" {
		// When running with Traefik, the prefix is already stripped
		contextPath = "/api/v1/feedbacks"
	} else {
		// When running standalone, include the service prefix
		contextPath = service + "/api/v1/feedbacks"
	}

	// Define Feedback endpoints with basic required
	var authGroup = server.Group(contextPath)
	authGroup.GET("", handler.GetFeedbacks)
	authGroup.GET("/user/:id", handler.GetFeedbacksByUser)
	authGroup.GET("/:id", handler.GetFeedbackById)
	authGroup.POST("", handler.CreateFeedback)
	authGroup.PUT("", handler.UpdateFeedback)
	authGroup.DELETE("", handler.RemoveFeedback)
	authGroup.GET("/test-grpc/:id", handler.TestGrpcFeedback)
	authGroup.GET("/tourGuide/:id", handler.GetTourGuideFeedbacks)
}
