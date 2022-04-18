// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}
	// 解析 json 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if errs != nil {
		// 验证失败，返回 422 转台吗和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
