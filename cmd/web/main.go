package main

import (

	"github.com/808-not-found/tik_duck/cmd/web/service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	go service.RunMessageServer()
	h := server.Default(server.WithHostPorts("0.0.0.0:8080"), server.WithMaxRequestBodySize(20<<20)) //nolint:all

	InitRouter(h)
	h.Spin()
}
