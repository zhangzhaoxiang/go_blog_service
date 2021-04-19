package main

import (
	"go_blog_service/global"
	"go_blog_service/internal/model"
	"go_blog_service/internal/routers"
	"go_blog_service/pkg/logger"
	setting2 "go_blog_service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
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
// 读取初始化配置
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	//err = setupDBEngine()
	//if err != nil {
	//	log.Fatalf("init.setupDBEngine err: %v", err)
	//}
}

func main() {
	//fmt.Println(global.ServerSetting)
	//fmt.Println(global.AppSetting)
	//fmt.Println(global.DatabaseSetting)
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
