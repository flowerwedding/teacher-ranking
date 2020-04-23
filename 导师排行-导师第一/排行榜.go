package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ranking(c *gin.Context){
	if yesterday == today{
		var teacherSlice []Teacher
		db.Where("identity = ?","选手").Order("count desc").Find(&teacherSlice)

		for _, teachers := range teacherSlice {
			c.JSON(http.StatusOK,gin.H{"username":teachers.Username,"count":teachers.Count})
		}
	}else
	{
		mapp, err := redis.StringMap(red.Do("ZRANGE", "teachers",0, -1,"withscores"))
		if err != nil {
			fmt.Println(err)
			return
		}

		for co, va := range mapp {
			db.Model(Teacher{}).Where("username = ?",co).Update("count",va)
		}

		yesterday = today
		c.JSON(http.StatusOK,mapp)
	}
}