package initialize

import (
	"fmt"
	"go-ecommerce-backend/global"
)

func Run() {
	// Load configuration
	LoadConfig()
	fmt.Println("loading config data", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()
	r.Run(":8002")
}
