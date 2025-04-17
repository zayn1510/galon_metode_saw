package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/routers"
)

func main() {
	c := gin.Default()
	routers.RegisterRoutes(c)
	c.Run(":8080")
}
