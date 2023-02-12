package main

import (
	"log"

	useruser "github.com/808-not-found/tik_duck/kitex_gen/useruser/useruserservice"
)

func main() {
	svr := useruser.NewServer(new(UserUserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
