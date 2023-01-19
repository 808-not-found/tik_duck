package main

import (
	"github.com/808-not-found/tik_duck/cmd/web/service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	go service.RunMessageServer()
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	InitRouter(h)
	h.Spin()
}
