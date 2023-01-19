package main

import (
	"context"

	"github.com/808-not-found/tik_duck/kitex_gen/douyin_user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// User_Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_Register(ctx context.Context, req *douyin_user.DouyinUserRegisterRequest) (resp *douyin_user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// User_GetFeed implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_GetFeed(ctx context.Context, req *douyin_user.DouyinFeedRequest) (resp *douyin_user.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// User_Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_Login(ctx context.Context, req *douyin_user.DouyinUserLoginRequest) (resp *douyin_user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// User_Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_Info(ctx context.Context, req *douyin_user.DouyinUserInfoRequest) (resp *douyin_user.DouyinUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// User_PublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_PublishList(ctx context.Context, req *douyin_user.DouyinPublishListRequest) (resp *douyin_user.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// User_PublishAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_PublishAction(ctx context.Context, req *douyin_user.DouyinPublishActionRequest) (resp *douyin_user.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}
