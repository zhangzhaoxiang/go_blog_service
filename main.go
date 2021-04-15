package main

import (
	"go_blog_service/internal/routers"
	"net/http"
	"time"
)

/*
@Time    : 2021/4/14 9:00 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : main.go
@Software: GoLand
*/

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
