// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// panic("测试 panic")
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}
	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
	//c.JSON(http.StatusOK, gin.H{
	//	"exist": user.IsPhoneExist(request.Phone),
	//})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exits": user.IsEmailExist(request.Email),
	})
	//c.JSON(http.StatusOK, gin.H{
	//	"exits": user.IsEmailExist(request.Email),
	//})
}
