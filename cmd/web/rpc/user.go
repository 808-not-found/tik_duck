package rpc

import (
	"time"

	userservice "github.com/808-not-found/tik_duck/kitex_gen/user/userservice"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client // nolint: all

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond), // nolint:all
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}
