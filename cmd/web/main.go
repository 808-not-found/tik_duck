package main

import (
	"os"

	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/cmd/web/service"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

func main() {
	rpc.InitRPC()
	go service.RunMessageServer()
	os.Mkdir("public", os.ModePerm)
	h := server.Default(
		server.WithStreamBody(true), server.WithTransport(standard.NewTransporter),
		server.WithHostPorts("0.0.0.0:"+consts.WebServerPort),
		server.WithMaxRequestBodySize(consts.HTTPMaxBodySize))

	InitRouter(h)
	h.Spin()
}
