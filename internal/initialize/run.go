package initialize

func Run() {
	// Load configuration
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()
	r.Run("8002")
}
