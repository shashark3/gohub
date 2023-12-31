package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"gohub/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	//记录错误日志，出错时记error等级的错误
	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id":   id,
		"captcha_b64s": b64s,
	})
}

// SendUsingPhone 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

	//1.验证表单
	request := requests.VerifyCodePhoneRequest{}

	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	//2.发送SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	}
	response.Success(c)

}

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	request := requests.VerifyCodeEmailRequest{}

	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	if ok := verifycode.NewVerifyCode().SendEmail(request.Email); ok != nil {
		response.Abort500(c, "发送邮件失败")
	}
	response.Success(c)
}
