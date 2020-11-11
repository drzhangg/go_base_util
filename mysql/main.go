package main

import (
	"github.com/gin-gonic/gin"
	"go_base_util/mysql/handler"
)

func main() {
	r := gin.New()

	r.POST("/register", handler.Register)

	r.Run(":8081")
}
