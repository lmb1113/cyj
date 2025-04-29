package main

import (
	"cyj/server/controller"
	"cyj/server/middleware"
	"cyj/server/pkg/logging"
	"github.com/gin-gonic/gin"
)

const (
	serverPort = ":10033"
)

func main() {
	logging.Setup()
	server := gin.New()
	server.Use(middleware.Cors())
	server.Use(middleware.Version())
	server.Use(middleware.GinLogger(logging.Logger))
	server.Use(middleware.WithRequestId())
	v1Api := server.Group("/api")
	v1Api.GET("/config", controller.Config)
	v1Api.GET("/update", controller.Update)
	v1Api.GET("/download", controller.Download)
	server.Run(serverPort)
}
