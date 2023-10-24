package middlewares

import (
	"fmt"
	"gohub/app/models/user"
	"gohub/pkg/config"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从标头Authorization:Bearer xxxx中获取信息，并验证JWT的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		//JWT解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看%v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		//JWT解析成功，设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已经删除")
			return
		}

		//将用户信息存入gin.context里，后续auth包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
