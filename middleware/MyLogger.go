package middleware

import (
	"blog/utils"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 日志拦截
func Logger() gin.HandlerFunc {
	logger := utils.Log()

	return func(c *gin.Context) {

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime) / time.Millisecond

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.Header

		// 请求头
		header := c.Handler()

		// 协议
		proto := c.Request.Proto

		// 状态码
		statusCode := c.Writer.Status()

		// 请求的ip
		clientIP := c.ClientIP()

		// 错误
		err := c.Err()

		// 请求提
		body, _ := ioutil.ReadAll((c.Request.Body))

		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"reqMethod":    reqMethod,
			"req_uri":      reqUri,
			"header":       header,
			"proto":        proto,
			"err":          err,
			"body":         body,
		}).Info()
	}
}
