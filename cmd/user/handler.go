package main

import (
	"context"

	douyinuser "github.com/808-not-found/tik_duck/kitex_gen/douyinuser"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// User_Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_Register(ctx context.Context, req *douyinuser.DouyinUserRegisterRequest) (resp *douyinuser.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// User_GetFeed implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_GetFeed(ctx context.Context, req *douyinuser.DouyinFeedRequest) (resp *douyinuser.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// User_Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_Login(ctx context.Context, req *douyinuser.DouyinUserLoginRequest) (resp *douyinuser.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// User_Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_Info(ctx context.Context, req *douyinuser.DouyinUserInfoRequest) (resp *douyinuser.DouyinUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// User_PublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_PublishList(ctx context.Context, req *douyinuser.DouyinPublishListRequest) (resp *douyinuser.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// User_PublishAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_PublishAction(ctx context.Context, req *douyinuser.DouyinPublishActionRequest) (resp *douyinuser.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// User_Test implements the UserServiceImpl interface.
func (s *UserServiceImpl) User_Test(ctx context.Context, req *douyinuser.DouyinTestinfo) (resp *douyinuser.DouyinTestinfo, err error) {
	// TODO: Your code here...
	resp = &douyinuser.DouyinTestinfo{Testinfo: req.Testinfo}
	return
}
