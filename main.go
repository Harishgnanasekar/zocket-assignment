package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"product-system/api"
	"product-system/cache"
	"product-system/database"
	"product-system/logger"
	"product-system/queue"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	logger.InitLogger()
	database.InitDB()
	cache.InitRedis()
	queue.InitQueue()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api.RegisterRoutes(r)

	log.Printf("Starting server on %s", viper.GetString("SERVER_PORT"))
	r.Run(viper.GetString("SERVER_PORT"))
}