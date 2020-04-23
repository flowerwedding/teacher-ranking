package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
var red  redis.Conn

func init() {
	var err error

	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/dome7?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	if !db.HasTable(&Teacher{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Teacher{}).Error; err != nil {
			panic(err)
		}
	}

	red, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	//defer red.Close()
}

type Teacher struct {
	Username  string `gorm:"type:varchar(256);not null;"`
	Password  string `gorm:"type:varchar(256);not null;"`
	Identity  string `gorm:"type:varchar(256);not null;"`
	Count     int    `gorm:"type:varchar(255);not null;"`
	Today     string `gorm:"type:varchar(256);not null;"`
}