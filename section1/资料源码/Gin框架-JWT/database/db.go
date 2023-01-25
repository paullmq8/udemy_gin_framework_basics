package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {

	var err error
	Db, err = gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/go_gin")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
