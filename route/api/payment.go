package api

import (
	"os"
	"tourmate/payment-service/handler"

	"github.com/gin-gonic/gin"
)

func InitializePaymentHandlerRoute(server *gin.Engine, service string) {
	//Context path
	var contextPath string
	if os.Getenv("DOCKER_COMPOSE") == "true" {
		// When running with Traefik, the prefix is already stripped
		contextPath = "/api/v1/payments"
	} else {
		// When running standalone, include the service prefix
		contextPath = service + "/api/v1/payments"
	}

	// Define Payment endpoints with admin required
	var adminAuthGroup = server.Group(contextPath)
	adminAuthGroup.GET("", handler.GetAllPayments)
	adminAuthGroup.PUT("/update", handler.UpdatePayment)

	// Define Payment endpoints with basic required
	var authGroup = server.Group(contextPath)
	authGroup.GET("/customer/:id", handler.GetPaymentsByUser)
	authGroup.GET("/:id", handler.GetPaymentById)
	authGroup.POST("/create", handler.CreatePayment)

	var norGroup = server.Group(contextPath)
	norGroup.POST("/create-embedded-payment-link", handler.CreatePayosTransaction)

}
