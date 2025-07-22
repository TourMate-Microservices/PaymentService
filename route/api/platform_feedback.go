package api

// import (
// 	"tourmate/payment-service/handler"
// 	"tourmate/payment-service/utils/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func InitializePlatformFeedbackHandlerRoute(server *gin.Engine, port, service string) {
// 	//Context path
// 	var contextPath string = "api/v1/platform-feedbacks"

// 	// Define Feedback endpoints with basic required
// 	var authGroup = server.Group(contextPath, middleware.Authorize)
// 	authGroup.GET("", handler.GetPlatformFeedbacks)
// 	authGroup.GET("/user/:id", handler.GetFeedbacksByUser)
// 	authGroup.GET("/:id", handler.GetPlatformFeedbackById)
// 	authGroup.POST("", handler.CreatePlatformFeedback)
// 	authGroup.PUT("", handler.UpdatePlatformFeedback)
// }
