package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	// 初始化 Gin 实例
	r := gin.New()

	// 注册中间件
	bootstrap.SetupRoute(r)

	// 运行服务
	err := r.Run(":8000")
	if err != nil {
		// 错误处理， 端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
