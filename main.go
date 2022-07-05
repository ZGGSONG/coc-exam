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
	"os/exec"
	"runtime"
	"time"
)

//go:embed web/*
var webFS embed.FS

const (
	PORT = 50020
)

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}
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
	go func() {
		time.Sleep(500 * time.Millisecond)
		Open("http://127.0.0.1:50020")
	}()
	panic(r.Run(fmt.Sprintf("0.0.0.0:%v", PORT)))
}
