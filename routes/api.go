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

			lgc := new(auth.LoginController)

			authGroup.POST("/login/using-phone", lgc.LoginByPhone)

			authGroup.POST("/login/using-password", lgc.LoginByPassword)

			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			pc := new(auth.PasswordController)
			//重置密码
			authGroup.POST("/password-reset/using-phone", pc.ResetByphone)

		}
	}
}
