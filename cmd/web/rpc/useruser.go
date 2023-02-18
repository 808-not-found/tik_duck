package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
	useruserservice "github.com/808-not-found/tik_duck/kitex_gen/useruser/useruserservice"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userUserClient useruserservice.Client // nolint: all

func initUserUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := useruserservice.NewClient(
		consts.UserServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond), // nolint:all
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userUserClient = c
}

func UserRelationAction(
	ctx context.Context, req *useruser.RelationActionRequest,
) (*useruser.RelationActionResponse, error) {
	resp, err := userUserClient.UserRelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserRelationFollowList(
	ctx context.Context, req *useruser.RelationFollowListRequest,
) (*useruser.RelationFollowListResponse, error) {
	resp, err := userUserClient.UserRelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserRelationFollowerList(
	ctx context.Context, req *useruser.RelationFollowerListRequest,
) (*useruser.RelationFollowerListResponse, error) {
	resp, err := userUserClient.UserRelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

func UserRelationFriendList(
	ctx context.Context, req *useruser.RelationFriendListRequest,
) (*useruser.RelationFriendListResponse, error) {
	resp, err := userUserClient.UserRelationFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg) // nolint: all
	}
	return resp, err
}

/*
service UserUserService {
    // 关系操作
    RelationActionResponse UserRelationAction(1:RelationActionRequest Req)
    // 关注列表
    RelationFollowListResponse UserRelationFollowList(1:RelationFollowListRequest Req)
    // 粉丝列表
    RelationFollowerListResponse UserRelationFollowerList(1:RelationFollowerListRequest Req)
    // 好友列表
    RelationFriendListResponse UserRelationFriendList(1:RelationFriendListRequest Req)
    // 消息操作暂且不实现
}
*/
