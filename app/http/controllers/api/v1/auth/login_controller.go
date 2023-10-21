package auth

import (
	v1 "gohub/app/http/controllers/api/v1"

	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {
	//1.验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	//2.尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Error(c, err, "账号不存在")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}

}
