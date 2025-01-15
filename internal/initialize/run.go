package initialize

import (
	"fmt"
	"go-ecommerce-backend/global"

	"go.uber.org/zap"
)

func Run() {
	// Load configuration
	LoadConfig()
	fmt.Println("loading config data", global.Config.Mysql.Username)
	fmt.Println("loading config data", global.Config.Logger)
	InitLogger()
	global.Logger.Info("config log ok ", zap.String("oke", "success"))
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()
	r.Run(":8002")
}
