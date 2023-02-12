package main

import (
	"log"

	userplat "github.com/808-not-found/tik_duck/kitex_gen/userplat/userplatservice"
)

func main() {
	svr := userplat.NewServer(new(UserPlatServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
