package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Init will create a db instance
func Init(user, password, hostname, port, dbName string) {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, hostname, port, dbName))
	if err != nil {
		panic(err)
	}
}
