package system

import (
	"coc-question-bank/router"
	"github.com/gin-gonic/gin"
)

func InitGin() {
	//初始化Gin
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		return
	}
	r = router.CollectRoute(r)

	panic(r.Run("0.0.0.0:8080"))
}
