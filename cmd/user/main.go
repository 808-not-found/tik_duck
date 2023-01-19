package main

import (
	"log"

	douyinuser "github.com/808-not-found/tik_duck/kitex_gen/douyinuser/userservice"
)

func main() {
	svr := douyinuser.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
