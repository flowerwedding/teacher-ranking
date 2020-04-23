package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var count int
var yesterday string
var today = time.Now().Format("2006-01-02")

func Register (c *gin.Context){
	var user Teacher
	user.Username=c.PostForm("username")
	db.Model(&Teacher{}).Where("username = ?",user.Username).First(&user)
	if user.Password != ""{
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "用户已注册"})
		return
	}

	user = Teacher{
		Password : c.PostForm("password"),
		Username : c.PostForm("username"),
		Identity : c.PostForm("identity"),//用户或选手
		Count    : 0,
		Today    : time.Now().Format("2006-01-02"),
	}
	if err :=  db.Model(&Teacher{}).Create(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "数据库报错", "question": "所有参数都有值吗"})
		return
	}

	if user.Identity == "选手"{
		_, err := red.Do("ZADD", "teachers", 0, user.Username)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "注册成功!"})
}

func Login (c *gin.Context){
	var user Teacher
	user.Username=c.DefaultPostForm("username","游客")
	user.Password=c.DefaultPostForm("password","")
	db.Model(&Teacher{}).Where(&Teacher{Username: user.Username, Password: user.Password}).First(&user)
	if user.Identity == "" {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "密码错误"})
		return
	}

	if user.Today != today{
		count = 3
		db.Model(Teacher{}).Where("username = ?",user.Username).Update("today",time.Now().Format("2006-01-02"))
	}

	singedToken,err:= RunToken(user.Username)
	if err != nil{
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "登录成功!","token":singedToken})
}