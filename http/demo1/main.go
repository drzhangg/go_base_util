package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.New()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"hello":"world"})
	})

	r.Run(":8888")

	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	request
	//})
}
