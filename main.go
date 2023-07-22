package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github/adekang/gin-demo/common"
	_ "github/adekang/gin-demo/docs"
	"net/http"
)

// @title Test Swagger Gin Demo API
// @version 1.0
// @description This is a sample example

// @contact.name   adekang
// @contact.url    https://github.com/adekang
// @contact.email  adekang@163.com

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	InitConfig()
	common.InitDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 8080
}

func InitConfig() {
	//workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	//viper.AddConfigPath(workDir + "/config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
