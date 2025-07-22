package utils

import (
	"log"
	"os"
)

func GetLogConfig() *log.Logger {
	var logger *log.Logger = log.New(os.Stdout, "[ERROR] ", log.LstdFlags)
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	return logger
}
