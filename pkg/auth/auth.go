package auth

//package auth 授权相关逻辑

import (
	"errors"
	"gohub/app/models/user"
	"gohub/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}

// CurrentUser 从gin.context中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}

	}
	return userModel
}

// CurrentUID 从gin.Context中获取当前登录用户ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
