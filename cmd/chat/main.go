package main

import (
	"log"

	chat "github.com/808-not-found/tik_duck/kitex_gen/chat/chatservice"
)

func main() {
	svr := chat.NewServer(new(ChatServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
