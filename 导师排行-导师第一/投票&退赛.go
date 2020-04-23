package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Exit (c *gin.Context) {
	username,_:=c.Get("username")

	_, err := red.Do("ZREM", "teachers", username)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Where("username=?",username).Delete(Teacher{})

	c.JSON(http.StatusOK,gin.H{"code":10000,"message":"退赛成功"})
}

func Vote(c *gin.Context){
	teacher:=c.PostForm("teacher")
	username,_:=c.Get("username")
    countant,_:=c.Get("count")

	if countant.(int) <= 0{
		c.JSON(http.StatusOK,gin.H{"code":10000,"message":"今天已投三票"})
		return
	}

	count--
	_, err1 := red.Do("ZINCRBY", "teachers", 1, teacher)
	_, err2 := red.Do("ZINCRBY", "teachers", 1, "游琎")
	if err1 != nil&&err2 != nil {
		fmt.Println(err1,err2)
		return
	}
	singedToken,err := RunToken(username.(string))
	if err != nil{
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK,gin.H{"code":10000,"message":"投票成功","token":singedToken})
}