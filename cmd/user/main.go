package main

import (
	"log"
	"net"

	server "github.com/cloudwego/kitex/server"

	user "github.com/808-not-found/tik_duck/kitex_gen/user/userservice"
)

func main() {
	addr,_ := net.ResolveTCPAddr("tcp" , "127.0.0.1:10001")
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

