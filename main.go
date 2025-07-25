package main

import "tourmate/payment-service/cmd"

// @title Tourmate - Payment Service API
// @version 1.0
// @description API for Payment service
// @host localhost:8080
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
