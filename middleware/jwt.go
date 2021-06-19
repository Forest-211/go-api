package middleware

import (
	"blog/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TokenData struct {
	Username string
	Password string
	Key      string
}

func (token *TokenData) SetToken() string {
	return utils.MD5(token.Username + token.Password + token.Key)
}
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		// token 为空
		if token == "" {
			res := &utils.Response{
				Code: 1100,
				Msg:  "请求未携带token，无权访问！",
			}
			res.Json(c)
			c.Abort()
			return
		}
		logrus.Debug("get token:", token)

		data, found := utils.Cache.Get(token)
		if found {
			res := &utils.Response{
				Code: 1100,
				Msg:  "token找不到或者已过期！",
			}
			res.Json(c)
			c.Abort()
			return
		}

		tokenData := data.(*TokenData)

		// 校验token是否一致
		isEqual := token != tokenData.SetToken()

		if isEqual {
			res := &utils.Response{
				Code: 1100,
				Msg:  "token认证出错！",
			}
			res.Json(c)
			c.Abort()
			return
		}

		c.Set("token", data)

		c.Next()
	}
}

// TODO: 返回参数公共封装
// func returnTip(message string, c *gin.Context) {
// 	res := &utils.Response{
// 		Code: 1100,
// 		Msg: message,
// 	}
// 	res.Json(c)
// 	c.Abort()
// 	return
// }
