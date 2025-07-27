package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tourmate/payment-service/constant/env"
	payment_env "tourmate/payment-service/constant/env/payment"
	"tourmate/payment-service/constant/noti"
	payment_method "tourmate/payment-service/constant/payment_method"
	"tourmate/payment-service/docs"
	api "tourmate/payment-service/route/api"
	grpc "tourmate/payment-service/route/gRPC"

	_ "tourmate/payment-service/docs"

	"github.com/gin-gonic/gin"
	"github.com/payOSHQ/payos-lib-golang"
	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
)

func setupApiRoutes(logger *log.Logger, service string) {
	// Initialize gin server for API
	var server = gin.Default()

	// Config CORS for requests
	corsConfig(server)

	// Get API port
	var apiPort string = os.Getenv(env.API_PORT)

	// Set up swagger FIRST (before any auth middleware)
	setupSwagger(server, service, apiPort)

	// Feedback API endpoints
	api.InitializeFeedbackHandlerRoute(server, service)

	// Payment API endpoints
	api.InitializePaymentHandlerRoute(server, service)

	// Default URL
	server.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/%s/swagger/index.html#", service))
	})

	// Run server
	if err := server.Run(":" + apiPort); err != nil {
		logger.Println("Error run service - " + err.Error())
	}
}

func setupSwagger(server *gin.Engine, service, port string) {
	//Configure swagger info
	docs.SwaggerInfo.Title = "Tourmate - Payment Service API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// If running in Docker Compose, set host to 'localhost' (no port), else use 'localhost:port'
	if os.Getenv("DOCKER_COMPOSE") == "true" {
		docs.SwaggerInfo.Host = "localhost"
	} else {
		docs.SwaggerInfo.Host = "localhost:" + port
	}

	//Add swagger route
	server.GET("/"+service+"/swagger/*any", gin_swagger.WrapHandler(swagger_files.Handler))
}

func setupGrpc(logger *log.Logger, service string) {
	// Initialize gRPC server
	grpc.InitializeGRPCRoute(logger, service)

	// Dial on gRPC servers

}

func setupPayments(logger *log.Logger) {
	// Payos
	if err := payos.Key(os.Getenv(payment_env.PAYOS_CLIENT_ID), os.Getenv(payment_env.PAYOS_API_KEY), os.Getenv(payment_env.PAYOS_CHECKSUM_KEY)); err != nil {
		logger.Println(fmt.Sprintf(noti.PAYMENT_INIT_ENV_ERR_MSG, payment_method.PAYOS) + err.Error())
	}
}
