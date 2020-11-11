package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id   int
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:root123456@tcp(47.103.9.218:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("err1:", err)
	}

	//defer Db.Close()
	Db = db
	Db.SingularTable(true)

}

func main() {
	r := gin.New()
	r.GET("/test", Test)
	r.Run(":8081")

}

func Test(c *gin.Context) {
	u1 := User{
		Name: "tom2",
		Age:  14,
	}
	err := u1.AddUser()
	fmt.Println(err)

	var uc = make(chan User)
	uc <- u1
	for i := 0; i < 5; i++ {
		go func() {
			u2 := <- uc
			u2.AddUser()
		}()
	}
}

func (u User) AddUser() error {
	err := Db.Table("user").Create(&u).Error
	return err
}
