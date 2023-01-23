package main

import (
	"context"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(
	ctx context.Context,
	req *user.UserRegisterRequest,
) (resp *user.UserRegisterResponse, err error) {
	//resp = new(user.FeedResponse)

	return
}

// 视频流
// UserGetFeed implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserGetFeed(
	ctx context.Context,
	req *user.FeedRequest,
) (resp *user.FeedResponse, err error) {
	resp = new(user.FeedResponse)

	statusCode, statusMsg, videoList, nextTime, err := userservice.UserGetFeedService(ctx, req)

	// if err = req.IsValid(); err != nil {
	// 	resp.StatusCode = 1001
	// 	return resp, nil
	// }
	if err != nil {
		return resp, err
	}
	resp.StatusCode = statusCode
	resp.StatusMsg = &statusMsg
	resp.VideoList = videoList
	resp.NextTime = &nextTime
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(
	ctx context.Context,
	req *user.UserLoginRequest,
) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(
	ctx context.Context,
	req *user.UserRequest,
) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// UserPublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserPublishList(
	ctx context.Context,
	req *user.PublishListRequest,
) (resp *user.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// UserPublishAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserPublishAction(
	ctx context.Context,
	req *user.PublishActionRequest,
) (resp *user.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// UserTest implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserTest(
	ctx context.Context,
	req *user.Testinfo,
) (resp *user.Testinfo, err error) {
	// TODO: Your code here...
	resp = &user.Testinfo{Testinfo: req.Testinfo}
	return
}
