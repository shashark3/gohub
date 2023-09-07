package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIroutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Hello": "world!",
			})
		})
	}
}
