package main

import (
	"boxDemo/drivers"
	"boxDemo/route"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	httpServer := gin.Default()
	route.RegisterRoutes(httpServer)
	defer drivers.MysqlDb.Close()
	_ = httpServer.Run(":8001")
}
