package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// 初始化 Gin 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		// 以 json 格式响应
		c.JSON(http.StatusOK, gin.H{"hello": "world!"})
	})

	// 处理 404 请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 html 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})

	// 运行服务
	r.Run(":8000")
}
