package system

import (
	"coc-question-bank/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	PORT = 50020
)

func InitGin() {
	//初始化Gin
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		return
	}

	r.LoadHTMLFiles("./web/index.html")
	r.StaticFile("/favicon.ico", "./web/favicon.ico")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "COC搜题",
			"icon":  "/favicon.ico",
		})
	})
	r = router.CollectRoute(r)
	fmt.Println(fmt.Sprintf("open http://127.0.0.1:%v", PORT))
	panic(r.Run(fmt.Sprintf("0.0.0.0:%v", PORT)))
}
