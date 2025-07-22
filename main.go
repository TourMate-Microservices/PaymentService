package main

import "tourmate/payment-service/cmd"

// @title Tourmate - Feedback Service API
// @version 1.0
// @description API for Feedback service
// @host localhost:8080
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
