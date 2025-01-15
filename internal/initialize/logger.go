package initialize

import (
	"go-ecommerce-backend/global"
	"go-ecommerce-backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}