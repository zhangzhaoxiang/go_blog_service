package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
@Time    : 2021/4/14 9:00 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : main.go
@Software: GoLand
*/

func main() {
	fmt.Println(111)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
