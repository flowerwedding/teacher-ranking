package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Bang struct {
	jwt.StandardClaims
	Username     string
	Count        int
}

func RunToken(username string) (string,error) {
	claims := &Bang{
		Username: username,
		Count   : count,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(3600)).Unix()
	var err error
	singedToken, err:= CreateToken(claims)
	if err != nil {
		return "",err
	}
	return singedToken,nil
}

func Middle(c *gin.Context) {
	auth:= c.GetHeader("Authorization")
	if len(auth)<7 {
		c.Abort()
		return
	}
	token, err := CheckAction(auth[:])
	if err != nil {
		fmt.Println(err.Error())
		c.Abort()
		return
	}
	c.Set("username", token.Username)
	c.Set("count", token.Count)
	c.Next()
	return
}

func CreateToken(claims *Bang) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}


func CheckAction(strToken string) (*Bang, error) {
	token, err := jwt.ParseWithClaims(strToken, &Bang{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Bang)
	if !ok {
		return nil, err
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	return claims, nil
}