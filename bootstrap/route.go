package bootstrap

//程序初始化
import (
	"gohub/app/http/middlewares"
	"gohub/pkg/response"
	"gohub/routes"
	"strings"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func SetUpRoute(router *gin.Engine) {
	//注册中间件
	registerGlobalMiddleWare(router)

	//注册API路由 GET POST ........
	routes.RegisterAPIroutes(router)

	//配置404路由
	setup404Handler(router)

}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(middlewares.Logger(), middlewares.Recovery())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			response.Abort404(ctx, "页面走丢了")
		} else {
			response.Abort404(ctx, "路由未定义")
		}
	})

}
