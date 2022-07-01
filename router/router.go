package router

import (
	"coc-question-bank/api"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/ping", api.Ping)
	r.GET("/search", api.Search)
	return r
}
