package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	//配置初始化，依赖命令行 --env参数
	var env string
	flag.StringVar(&env, "env", "", "加载.env文件，如 --env=testing加载的是 .env.testing文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Gin 实例
	r := gin.New()

	//初始化DB
	bootstrap.SetupDB()

	// 注册一个路由
	bootstrap.SetUpRoute(r)

	// 运行服务
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
