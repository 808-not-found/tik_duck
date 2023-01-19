package main

import (
	"log"

	"github.com/808-not-found/tik_duck/kitex_gen/douyin_user"
)

func main() {
	svr := douyin_user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
