package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go_base_util/goroutine/demo3"
	"go_base_util/goroutine/handler"
	_ "net/http/pprof"
)

func main() {
	wm := demo3.NewWorkerManager(10)
	wm.StartWorkerPool()

}

func Test(c *gin.Context) {
	var u handler.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(500, gin.H{
			"data": nil,
			"msg":  "bind error",
		})
		return
	}

	var uc = make(chan handler.User)
	uc <- u
	go func() {
		u2 := <-uc
		u2.AddUser()
	}()

	c.JSON(200, gin.H{
		"data": "ok",
	})
}
