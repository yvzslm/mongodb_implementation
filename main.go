package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-mongodb-implementation/api"
	"github.com/go-mongodb-implementation/mongodb"
	"github.com/go-mongodb-implementation/pkg/config"
	"github.com/go-mongodb-implementation/repository"
)

const (
	TIMEOUT       = 10
	DATABASE_NAME = "users"
)

var (
	cfg *config.AppConfig
)

func main() {
	cfg = config.GetConfig()

	var mongoDbClient, err = mongodb.NewClient(context.TODO(), cfg.MongoDB, TIMEOUT*time.Second)
	if err != nil {
		panic("MongoDB connection error. " + err.Error())
	}

	var userRepository = repository.NewUserRepository(mongoDbClient, DATABASE_NAME)

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost, 127.0.0.1"})
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", func(ctx *gin.Context) {
			api.GetUsers(ctx, userRepository)
		})
		v1.POST("/users", func(ctx *gin.Context) {
			api.InsertUser(ctx, userRepository)
		})
	}

	router.Run(":1453")
}
