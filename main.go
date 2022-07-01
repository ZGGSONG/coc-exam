package main

import (
	"coc-question-bank/system"
	"log"
)

func main() {
	_, err := system.InitDB()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Connecting to database successfully...")

}
