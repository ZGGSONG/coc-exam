package main

import (
	"coc-question-bank/model"
	"coc-question-bank/system"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	FILEPATH = "/Users/song/projects/go/coc-question-bank/util/coc.txt"
)

//导入TXT至数据库
func main() {
	//db
	db, err := system.InitDB()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Connecting to database successfully...")

	//读取文件
	bytes, err := readFile(FILEPATH)
	if err != nil {
		log.Fatalln(err)
		return
	}
	//log.Println(string(bytes))
	var mType, mQuestion, mOptions, mAnswer string
	msg := strings.Split(string(bytes), "\r\n")
	for _, value := range msg {
		//无效数据
		if value == "" || value == "\n" {
			continue
		}
		//fmt.Println(value)
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
			mQuestion = value
		} else if strings.Contains("ABCDEFGHIJK", firstCharacter) { //选项
			if strings.Contains(value, "、") {
				mOptions += value + " "
			}
		} else if firstCharacter == "提" || firstCharacter == "正" { //答案
			comma := strings.Index(value, "：")
			pos := strings.Index(value[comma:], "")
			mAnswer = value[comma+pos+3:] //取出答案
			//fmt.Println(mType + " " + mQuestion + " " + mOptions + " " + mAnswer)
			db.Create(&model.Subject{Type: mType, Question: mQuestion, Options: mOptions, Answer: mAnswer, UpdateTime: time.Now()})
			//清空变量
			mOptions = ""
			mAnswer = ""
			mQuestion = ""
		}
	}
	fmt.Println("插入完成...")
}

func readFile(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("read file fail" + err.Error())
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.New("read file fail" + err.Error())
	}
	return fd, nil
}
