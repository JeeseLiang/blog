package main

import (
	"blog/global"
	"blog/internal/routers"
	"blog/pkg"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupSetting() error {
	s, err := pkg.NewSetting()
	if err != nil {
		log.Fatalf("Init failed : %v\n", err)
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		log.Fatalf("Init failed : %v\n", err)
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		log.Fatalf("Init failed : %v\n", err)
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		log.Fatalf("Init failed : %v\n", err)
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func init() { //初始化
	// 配置初始化工作
	setupSetting()

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
