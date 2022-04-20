package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// Package auth 处理用户身份认证相关逻辑
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			// 发送短信验证码
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			// 发送邮件验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 账号密码登录，账号支持手机号，email，用户名
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
			// token 过期，刷新token
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)
		}
	}
}
