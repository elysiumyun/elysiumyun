package middleware

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func GormLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
			SlowThreshold:             200 * time.Millisecond,
		},
	)
}
