package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "")
	if err != nil {
		fmt.Println("err1:", err)
	}

	//defer Db.Close()
	Db = db
	Db.SingularTable(true)

}
