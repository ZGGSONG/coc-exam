package main

import (
	"coc-question-bank/global"
	"coc-question-bank/router"
	"coc-question-bank/system"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed web/*
var webFS embed.FS

const (
	PORT = 50020
)

func main() {
	db, err := system.InitDB()
	if err != nil {
		log.Fatalln(err)
		return
	}
	global.GLO_DB = db

	r := gin.Default()
	err = r.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		return
	}

	temple := template.Must(template.New("").ParseFS(webFS, "web/*.html"))
	r.SetHTMLTemplate(temple)
	fe, _ := fs.Sub(webFS, "web")
	r.StaticFS("/static", http.FS(fe))
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "COC搜题",
			"icon":  "/static/favicon.ico",
		})
	})
	r = router.CollectRoute(r)
	fmt.Println(fmt.Sprintf("open http://127.0.0.1:%v", PORT))
	panic(r.Run(fmt.Sprintf("0.0.0.0:%v", PORT)))
}
