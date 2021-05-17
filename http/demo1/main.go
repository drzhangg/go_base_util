package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.New()

	var a []int

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"hello":a})
	})

	r.Run(":8888")

	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	request
	//})
}
