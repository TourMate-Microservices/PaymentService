package cmd

import (
	"os"
	"tourmate/payment-service/constant/env"
	"tourmate/payment-service/utils"
)

// Execute server application
func Execute() {
	// Initialize logger config
	var logger = utils.GetLogConfig()

	// Load env
	loadEnv(logger)

	// Setup payments
	setupPayments(logger)

	// Get service name
	var service string = os.Getenv(env.SERVICE_NAME)

	// Setup gRPC routes
	go setupGrpcRoutes(logger, service)

	// Setup API routes
	setupApiRoutes(logger, service)
}
