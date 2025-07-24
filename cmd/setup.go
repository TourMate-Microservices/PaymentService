package cmd

import (
	"log"
	"os"
	"tourmate/payment-service/constant/env"
	"tourmate/payment-service/docs"
	api "tourmate/payment-service/route/api"

	_ "tourmate/payment-service/docs"

	"github.com/gin-gonic/gin"
	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
)

func setupApiRoutes(logger *log.Logger) {
	// Initialize gin server for API
	var server = gin.Default()

	// Config CORS for requests
	corsConfig(server)

	// Get API port
	var apiPort string = os.Getenv(env.API_PORT)

	// Set up swagger
	setupSwagger(server, apiPort)

	// Get service name
	var service string = os.Getenv(env.SERVICE_NAME)

	logger.Println(service)

	// Feedback API endpoints
	api.InitializeFeedbackHandlerRoute(server, apiPort, service)

	// Payment API endpoints
	api.InitializePaymentHandlerRoute(server, apiPort, service)

	// Run server
	if err := server.Run(":" + apiPort); err != nil {
		logger.Println("Error run service - " + err.Error())
	}
}

func setupSwagger(server *gin.Engine, port string) {
	//Configure swagger info
	docs.SwaggerInfo.Title = "Tourmate - Payment Service API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = "localhost:" + port

	//Add swagger route
	server.GET("/payment-service/swagger/*any", gin_swagger.WrapHandler(swagger_files.Handler))
}

func setupGrpcRoutes(logger *log.Logger) {

}
