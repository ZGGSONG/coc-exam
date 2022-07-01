package api

import (
	"coc-question-bank/global"
	"coc-question-bank/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Pong",
	})
}

func Search(ctx *gin.Context) {
	var model []model.Subject
	var que = ctx.Query("question")
	fmt.Println(que)
	tmp := fmt.Sprintf("%v%v%v", "%", que, "%")
	global.GLO_DB.Where("Question LIKE ?", tmp).Find(&model)
	switch len(model) {
	case 0:
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "No Records Found...",
			"data": "",
		})
	case 1:
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Found One Record...",
			"data": model[0],
		})
	default:
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Found Multiple Records...",
			"data": model,
		})
	}
}
