package api

import (
	"tourmate/payment-service/handler"
	"tourmate/payment-service/utils/middleware"

	"github.com/gin-gonic/gin"
)

func InitializePaymentHandlerRoute(server *gin.Engine, port, service string) {
	//Context path
	var contextPath string = service + "/api/v1/payments"

	// Define Payment endpoints with admin required
	var adminAuthGroup = server.Group(contextPath)
	adminAuthGroup.GET("", handler.GetAllPayments)
	adminAuthGroup.PUT("/update", handler.UpdatePayment)

	// Define Payment endpoints with basic required
	var authGroup = server.Group(contextPath)
	authGroup.GET("/customer/:id", handler.GetPaymentsByUser)
	authGroup.GET("/:id", handler.GetPaymentById)

	var norGroup = server.Group(contextPath)
	norGroup.GET("/callback-success/:id", handler.CallbackPaymentSuccess)
	norGroup.GET("/callback-cancel/:id", handler.CallbackPaymentCancel)
}
