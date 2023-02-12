package main

import (
	"github.com/808-not-found/tik_duck/cmd/web/rpc"
	"github.com/808-not-found/tik_duck/cmd/web/service"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	rpc.InitRPC()
	go service.RunMessageServer()
	h := server.Default(server.WithHostPorts("0.0.0.0:"+consts.WebServerPort),
		server.WithMaxRequestBodySize(consts.HTTPMaxBodySize))

	InitRouter(h)
	h.Spin()
}
