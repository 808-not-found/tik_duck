package main

import (
	"context"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// User注册.
func (s *UserServiceImpl) UserRegister(
	ctx context.Context,
	req *user.UserRegisterRequest,
) (resp *user.UserRegisterResponse, err error) {
	// 1: i32 StatusCode //状态码，0-成功，其他值失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: i64 UserId //用户id
	// 4: string Token //用户鉴权token

	// 生成回应结构体
	resp = new(user.UserRegisterResponse)
	// 校验参数
	err = req.IsValid()
	if err != nil {
		return resp, err
	}
	// 实现逻辑
	StatusCode, StatusMsg, UserID, Token, err := userservice.UserRegisterService(ctx, req)
	if err != nil {
		return resp, err
	}
	resp.StatusCode = StatusCode
	resp.StatusMsg = &StatusMsg
	resp.UserId = UserID
	resp.Token = Token
	// 返回结构体
	return resp, nil
}

// 视频流接口
// UserGetFeed implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserGetFeed(
	ctx context.Context,
	req *user.FeedRequest,
) (resp *user.FeedResponse, err error) {
	// 1: i32 StatusCode //状态码，0-成功，其他值失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: list<Video> VideoList //视频列表
	// 4: optional i64 NextTime //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time

	// 创建回应结构体
	resp = new(user.FeedResponse)

	// 判断请求是否合法
	if err = req.IsValid(); err != nil {
		return resp, err
	}

	statusCode, statusMsg, videoList, nextTime, err := userservice.UserGetFeedService(ctx, req)
	if err != nil {
		resp.StatusCode = 1101
		return resp, err
	}
	resp.StatusCode = statusCode
	resp.StatusMsg = &statusMsg
	resp.VideoList = videoList
	resp.NextTime = &nextTime

	return resp, nil
}

// User登陆.
func (s *UserServiceImpl) UserLogin(
	ctx context.Context,
	req *user.UserLoginRequest,
) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	// 1: i32 StatusCode //状态码，0-成功，其他值失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: i64 UserId //用户id
	// 4: string Token //用户鉴权token

	// 生成回应结构体
	resp = new(user.UserLoginResponse)
	// 校验参数
	err = req.IsValid()
	if err != nil {
		return resp, err
	}
	// 实现逻辑
	StatusCode, StatusMsg, UserID, Token, err := userservice.UserLoginService(ctx, req)
	if err != nil {
		return resp, err
	}
	resp.StatusCode = StatusCode
	resp.StatusMsg = &StatusMsg
	resp.UserId = UserID
	resp.Token = Token
	// 返回结构体
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(
	ctx context.Context,
	req *user.UserRequest,
) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	// 1: i32 StatusCode //状态码，0-成功，其他值-失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: User User //用户信息
	// 生成回应结构体
	resp = new(user.UserResponse)
	// 校验参数
	err = req.IsValid()
	if err != nil {
		return resp, err
	}
	// 实现逻辑
	StatusCode, StatusMsg, UserInfo, err := userservice.UserInfoService(ctx, req)
	if err != nil {
		return resp, err
	}
	resp.StatusCode = StatusCode
	resp.StatusMsg = &StatusMsg
	resp.User = UserInfo
	// 返回结构体
	return resp, nil
}

// 视频列表接口
// UserPublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserPublishList(
	ctx context.Context,
	req *user.PublishListRequest,
) (resp *user.PublishListResponse, err error) {
	// TODO: Your code here...
	// 1: i32 StatusCode //状态码，0-成功，其他值-失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: list<Video> VideoList //用户发布的视频列表

	// 生成返回结构体
	// resp = new(user.PublishListResponse)

	// // 判断请求是否合法
	// if err = req.IsValid(); err != nil {
	// 	resp.StatusCode = 1101
	// 	return resp, err
	// }

	// videoList, err = userservice.UserPublishList(req.UserId)

	return resp, nil
}

// 视频发布接口
// UserPublishAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserPublishAction(
	ctx context.Context,
	req *user.PublishActionRequest,
) (resp *user.PublishActionResponse, err error) {
	// TODO: Your code here...
	// 1: i32 StatusCode //状态码，0-成功，其他值-失败
	// 2: optional string StatusMsg //返回状态描述
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
