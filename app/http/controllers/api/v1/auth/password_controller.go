package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

// ResetByphone 用手机重置密码
func (pc *PasswordController) ResetByphone(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	}
	userModel.Password = request.Password
	userModel.Save()

	//操作成功返回信息
	response.Success(c)
}
