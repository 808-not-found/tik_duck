package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/808-not-found/tik_duck/kitex_gen/user"
	userservice "github.com/808-not-found/tik_duck/kitex_gen/user/userservice"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/middleware"
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
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
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

func GetFeed(ctx context.Context, req *user.FeedRequest) (*user.FeedResponse, error) {
	resp, err := userClient.UserGetFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserInfo(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserPublishList(ctx context.Context, req *user.PublishListRequest) (*user.PublishListResponse, error) {
	resp, err := userClient.UserPublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserPublishAction(ctx context.Context, req *user.PublishActionRequest) (*user.PublishActionResponse, error) {
	resp, err := userClient.UserPublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}
