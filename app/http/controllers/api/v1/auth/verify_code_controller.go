package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"gohub/pkg/verifycode"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 获取图形验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证吗
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
	//c.JSON(http.StatusOK, gin.H{
	//	"captcha_id":    id,
	//	"captcha_image": b64s,
	//})
}

// SendUsingPhone 发送短信验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 2. 发送 sms
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败～")
	} else {
		response.Success(c)
	}
}

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	// 2. 发送邮件验证码
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送 Email 验证码失败～")
	} else {
		response.Success(c)
	}
}
