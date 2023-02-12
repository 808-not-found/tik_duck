package main

import (
	useruser "github.com/808-not-found/tik_duck/IDLs/kitex_gen/useruser/useruserservice"
	"log"
)

func main() {
	svr := useruser.NewServer(new(UserUserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
