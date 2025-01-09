package initialize

import (
	"fmt"
	"go-ecommerce-backend/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("./config/")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	// read configuration
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}
	// read server configuration
	fmt.Println("Server.Port ::", v.GetInt("server.port"))
	fmt.Println("Server.Port ::", v.GetString("security.jwt.key"))
	// config structure

	if err := v.Unmarshal((&global.Config)); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

}
