package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

const (
	BASEURL = ""
)

func init() {
	db, err := gorm.Open("mysql", "root:root123456@tcp(47.103.9.218:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("err1:", err)
	}

	//defer Db.Close()
	Db = db
	Db.SingularTable(true)

}
