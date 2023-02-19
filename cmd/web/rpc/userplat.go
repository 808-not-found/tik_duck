package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/808-not-found/tik_duck/kitex_gen/userplat"
	userplatservice "github.com/808-not-found/tik_duck/kitex_gen/userplat/userplatservice"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userPlatClient userplatservice.Client // nolint: all

func initUserPlatRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userplatservice.NewClient(
		consts.UserPlatServiceName,
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
	userPlatClient = c
}

func UserFavoriteAction(
	ctx context.Context, req *userplat.FavoriteActionRequest,
) (*userplat.FavoriteActionResponse, error) {
	resp, err := userPlatClient.UserFavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserFavoriteList(ctx context.Context, req *userplat.FavoriteListRequest) (*userplat.FavoriteListResponse, error) {
	resp, err := userPlatClient.UserFavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserCommentAction(
	ctx context.Context, req *userplat.CommentActionRequest,
) (*userplat.CommentActionResponse, error) {
	resp, err := userPlatClient.UserCommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserCommentList(ctx context.Context, req *userplat.CommentListRequest) (*userplat.CommentListResponse, error) {
	resp, err := userPlatClient.UserCommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}
