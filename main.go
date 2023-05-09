package main

import (
	"backend_gin/common"
	"backend_gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = routes.CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run())
	}

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
