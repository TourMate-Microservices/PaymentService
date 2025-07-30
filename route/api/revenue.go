package api

import (
	"os"
	"tourmate/payment-service/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRevenueHandlerRoute(server *gin.Engine, service string) {
	//Context path
	var contextPath string
	if os.Getenv("DOCKER_COMPOSE") == "true" {
		// When running with Traefik, the prefix is already stripped
		contextPath = "/api/v1/revenues"
	} else {
		// When running standalone, include the service prefix
		contextPath = service + "/api/v1/revenues"
	}

	// Define Feedback endpoints with basic required
	var authGroup = server.Group(contextPath)
	authGroup.GET("", handler.GetRevenues)
	authGroup.GET("/monthly/:id", handler.GetMonthlyRevenue)
	authGroup.GET("/growth/:id", handler.GetGrowthPercentage)
	authGroup.GET("/stats/:id", handler.GetRevenueStats)
	authGroup.GET("/:id", handler.GetRevenue)
	authGroup.POST("", handler.CreateRevenue)
	authGroup.PUT("/:id", handler.UpdateRevenue)
	authGroup.DELETE("/:id", handler.RemoveRevenue)
}
