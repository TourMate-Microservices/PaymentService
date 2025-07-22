package cmd

import (
	"fmt"
	"log"
	"time"
	"tourmate/payment-service/constant/noti"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Load .env file
func loadEnv(logger *log.Logger) {
	if err := godotenv.Load(); err != nil {
		logger.Fatal(fmt.Sprintf(noti.ENV_LOAD_ERR_MSG, "") + err.Error())
	}
}

// Enable CORS
func corsConfig(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins, or specify ["http://example.com"]
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
