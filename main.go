package main

import (
	"blog/global"
	"blog/internal/model"
	"blog/internal/routers"
	"blog/pkg"
	"blog/pkg/logger"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func setupSetting() error {
	s, err := pkg.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {

		return err
	}

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDB() (err error) {
	global.DBEngine, err = model.NewDBEngine()
	if err != nil {
		log.Fatalf("NewDBEngine failed : %v\n", err)
	}
	return
}
func setupLogger() (err error) {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger("", log.LstdFlags, &lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   60,
		MaxAge:    10,
		LocalTime: true,
	}).WithCaller(2)
	return nil
}
func init() { //初始化

	// 配置初始化工作
	err := setupSetting()
	if err != nil {
		log.Fatalf("Setup Setting failed : %v\n", err)
	}
	// 数据库初始化
	err = setupDB()
	if err != nil {
		log.Fatalf("Setup DB failed : %v\n", err)
	}
	// 日志初始化
	err = setupLogger()
	if err != nil {
		log.Fatalf("Setup Logger failed : %v\n", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.Port,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		MaxHeaderBytes: 1 << 20, // 1MB
		WriteTimeout:   global.ServerSetting.WriteTimeout,
	}
	s.ListenAndServe()
}
