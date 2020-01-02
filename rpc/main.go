package main

import (
	"boxDemo/route"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	route.RegisterRoutes(r)
	r.Run(":8082")
}
