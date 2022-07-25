package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	setupRedis()
	setupMongoDB()
	setupWebServer()
}

func setupRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func setupMongoDB() {
	err := mgm.SetDefaultConfig(nil, "game_server", options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		panic(err)
	}
}

func setupWebServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8000")
}
