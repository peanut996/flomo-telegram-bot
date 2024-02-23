package main

import (
	"log"
)

func init() {
	log.Println("flomo bot starting ...")
	if err := retrieveToken(); err != nil {
		log.Fatal(err)
		return
	}

}

func main() {
	bot, err := NewBot(TgBotToken)
	if err != nil {
		panic(err)
	}
	log.Println("flomo bot start done!")
	bot.Run()
}
