package main

import (
	"blog/conf"
	"blog/db"
	"blog/router"
	"blog/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Db.Close()

	log := utils.Log()

	gin.SetMode(conf.Conf.Server.Model)

	// 路由
	router := router.InitRouter()

	srv := &http.Server{
		Addr:    conf.Conf.Server.Address,
		Handler: router,
	}

	go func() {
		// 启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s \n", err)
		}
		log.Fatal("listen: %s \n", conf.Conf.Server.Address)
	}()

	// 监听服务停机、宕机后的相关打印
	quit := make(chan os.Signal)
	//监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
