package main
//除游琎导师外其余导师收到一票加一票，任何人收到一票游琎导师的票数都会增加

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	router.Use(cors.Default())

	router.POST("/Teacher/register",Register)//用户注册，选手参赛
	router.POST("/Teacher/login",Login)//用户,选手登录

	router.DELETE("/Teacher/exit",Middle,Exit)//选手退赛
	router.POST("/Teacher/vote",Middle,Vote)//用户投票

	router.POST("/Teacher/ranking",Ranking)//比赛排行,每日更新

	_ = router.Run(":8080")
}

