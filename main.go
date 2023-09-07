package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	r := gin.New()

	// 注册一个路由
	bootstrap.SetUpRoute(r)

	// 运行服务
	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
