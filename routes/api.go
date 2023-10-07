package routes

import (
	"gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAPIroutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)

			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)

			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			vcc := new(auth.VerifyCodeController)
			//显示验证码
			authGroup.POST("/verify_code/captcha", vcc.ShowCaptcha)

			authGroup.POST("/verify_code/phone", vcc.SendUsingPhone)

			authGroup.POST("/verify_code/email", vcc.SendUsingEmail)

		}
	}
}
