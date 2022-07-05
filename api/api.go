package api

import (
	"coc-question-bank/global"
	"coc-question-bank/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
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

func Put(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "失败: " + err.Error(),
			"data": "",
		})
		return
	}
	//读取文件
	open, err := file.Open()
	defer open.Close()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "失败: " + err.Error(),
			"data": "",
		})
		return
	}
	bytes, err := ioutil.ReadAll(open)
	var mType, mQuestion, mOptions, mAnswer string
	var usefulCount, uselessCount int64
	msg := strings.Split(string(bytes), "\r\n")
	for _, value := range msg {
		//无效数据
		if value == "" || value == "\n" {
			continue
		}
		if strings.Contains(value, "单选题") {
			mType = "single"
			continue
		}
		if strings.Contains(value, "多选题") {
			mType = "multiple"
			continue
		}
		if strings.Contains(value, "判断题") {
			mType = "judgement"
			continue
		}

		//获取第一个字符
		firstCharacter := regexp.MustCompile("^.").FindStringSubmatch(value)[0]
		if strings.Contains("1234567890", firstCharacter) { //题目
			com := strings.Index(value, ".")
			mQuestion = value[com+1:] //取出结果
		} else if strings.Contains("ABCDEFGHIJK", firstCharacter) { //选项
			if strings.Contains(value, "、") {
				mOptions += value + " "
			}
		} else if firstCharacter == "提" || firstCharacter == "正" { //答案
			comma := strings.Index(value, "：")
			pos := strings.Index(value[comma:], "")
			mAnswer = value[comma+pos+3:] //取出答案
			//fmt.Println(mType + " " + mQuestion + " " + mOptions + " " + mAnswer)
			var tmp model.Subject
			global.GLO_DB.Where("Question", mQuestion).First(&tmp)
			if tmp.ID == 0 {
				global.GLO_DB.Create(&model.Subject{Type: mType, Question: mQuestion, Options: mOptions, Answer: mAnswer, UpdateTime: time.Now()})
				usefulCount++
			} else {
				uselessCount++
			}
			//清空变量
			mOptions = ""
			mAnswer = ""
			mQuestion = ""
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf("插入%v条有效数据,过滤%v无效数据", usefulCount, uselessCount),
		"data": "",
	})
}
