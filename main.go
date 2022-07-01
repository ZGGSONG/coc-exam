package main

import (
	"coc-question-bank/global"
	"coc-question-bank/system"
	"log"
)

func main() {
	db, err := system.InitDB()
	if err != nil {
		log.Fatalln(err)
		return
	}
	global.GLO_DB = db
	log.Println("Connecting to database successfully...")

	system.InitGin()
	log.Println("Init Gin successfully...")
}
