package main

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/useruser/uuservice"
	useruser "github.com/808-not-found/tik_duck/kitex_gen/useruser"
)

// UserUserServiceImpl implements the last service interface defined in the IDL.
type UserUserServiceImpl struct{}

// 登录用户对其他用户进行关注或取消关注。
// UserRelationAction implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationAction(
	ctx context.Context,
	req *useruser.RelationActionRequest,
) (resp *useruser.RelationActionResponse, err error) {
	// TODO: Your code here...
	if err = req.IsValid(); err != nil {
		return resp, err
	}

	resp, err = uuservice.UserRelationActionService(ctx, req)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// 登录用户关注的所有用户列表。
// UserRelationFollowList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFollowList(
	ctx context.Context,
	req *useruser.RelationFollowListRequest,
) (resp *useruser.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	if err = req.IsValid(); err != nil {
		return resp, err
	}

	resp, err = uuservice.UserRelationFollowListService(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// 所有关注登录用户的粉丝列表。
// UserRelationFollowerList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFollowerList(
	ctx context.Context,
	req *useruser.RelationFollowerListRequest,
) (resp *useruser.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	if err = req.IsValid(); err != nil {
		return resp, err
	}

	resp, err = uuservice.UserRelationFollowerList(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// 所有关注登录用户的好友列表。
// UserRelationFriendList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFriendList(
	ctx context.Context,
	req *useruser.RelationFriendListRequest,
) (resp *useruser.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	if err = req.IsValid(); err != nil {
		return resp, err
	}

	resp, err = uuservice.UserRelationFriendListService(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
