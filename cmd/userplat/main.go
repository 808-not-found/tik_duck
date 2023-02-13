package main

import (
	"log"
	"net"

	db "github.com/808-not-found/tik_duck/cmd/userplat/dal/db"
	userplat "github.com/808-not-found/tik_duck/kitex_gen/userplat/userplatservice"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	server "github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:10001")
	svr := userplat.NewServer(new(UserPlatServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserPlatServiceName}), // server name
		server.WithServiceAddr(addr), server.WithRegistry(r))
	db.Init()
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
