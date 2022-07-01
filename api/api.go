package api

import (
	"coc-question-bank/global"
	"coc-question-bank/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Pong",
		"data": "",
	})
}

func Search(ctx *gin.Context) {
	var model []model.Subject
	var que = ctx.Query("question")
	if que == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusOK,
			"msg":  "Please Input your question...",
			"data": "",
		})
		return
	}
	tmp := fmt.Sprintf("%v%v%v", "%", que, "%")
	global.GLO_DB.Where("Question LIKE ?", tmp).Find(&model)
	switch len(model) {
	case 0:
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "No Records Found...",
			"data": "",
		})
	case 1: //搜出一个答案
		var answers string
		mOptions := strings.Split(model[0].Options, " ")
		for _, option := range mOptions {
			switch model[0].Type {
			case "multiple":
				ans := strings.Split(model[0].Answer, "、")
				for _, v := range ans {
					if strings.Contains(option, v) {
						answers += option + " "
					}
				}
			default:
				if strings.Contains(option, model[0].Answer) {
					answers = option
				}
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Successfully",
			"data": model,
		})
	default:
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Successfully",
			"data": model,
		})
	}
}
