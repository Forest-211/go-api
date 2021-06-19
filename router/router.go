package router

import (
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()

	// 宕机时可以恢复
	router.Use(gin.Recovery())

	// 静态文件
	router.Static("/static", "static")

	// 跨域中间件
	router.Use(middleware.Cors())

	register(router)

	// 日志中间件
	router.Use(middleware.Logger())



	return router
}

func register (router *gin.Engine) {
	
}
