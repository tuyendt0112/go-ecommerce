package global

import (
	"go-ecommerce-backend/pkg/logger"
	"go-ecommerce-backend/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
