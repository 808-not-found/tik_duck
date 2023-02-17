package main

import (
	"context"

	userplatservice "github.com/808-not-found/tik_duck/cmd/userplat/userplatService"
	userplat "github.com/808-not-found/tik_duck/kitex_gen/userplat"
)

// UserPlatServiceImpl implements the last service interface defined in the IDL.
type UserPlatServiceImpl struct{}

// UserFavoriteAction implements the UserPlatServiceImpl interface.
func (s *UserPlatServiceImpl) UserFavoriteAction(
	ctx context.Context,
	req *userplat.FavoriteActionRequest,
) (resp *userplat.FavoriteActionResponse, err error) {
	//1: i32 StatusCode           //状态码，0-成功，其他值失败
	//2: string StatusMsg         // 返回状态描述

	//生成回应结构体
	resp = new(userplat.FavoriteActionResponse)
	//校验参数
	err = req.IsValid()
	if err != nil {
		resp.StatusCode = 1201
		return resp, err
	}
	// 实现逻辑
	//StatusCode, StatusMsg
	resp, err = userplatservice.UserFavoriteActionService(ctx, req)
	if err != nil {
		resp.StatusCode = 1202
		return resp, err
	}

	return resp, nil

}

// UserFavoriteList implements the UserPlatServiceImpl interface.
func (s *UserPlatServiceImpl) UserFavoriteList(
	ctx context.Context,
	req *userplat.FavoriteListRequest,
) (resp *userplat.FavoriteListResponse, err error) {
	//1: i32 StatusCode //状态码，0-成功，其他值失败
	//2: optional string StatusMsg //返回状态描述
	//3: list<Video> VideoList //用户点赞视频列表

	//生成回应结构体
	resp = new(userplat.FavoriteListResponse)
	//校验参数
	err = req.IsValid()
	if err != nil {
		resp.StatusCode = 1201
		return resp, err
	}
	//实现逻辑
	//StatusCode, StatusMsg, VideoList
	resp, err = userplatservice.UserFavoriteListService(ctx, req)
	if err != nil {
		resp.StatusCode = 1202
		return resp, err
	}

	return resp, nil

}

// UserCommentAction implements the UserPlatServiceImpl interface.
func (s *UserPlatServiceImpl) UserCommentAction(
	ctx context.Context,
	req *userplat.CommentActionRequest,
) (resp *userplat.CommentActionResponse, err error) {
	//1: i32 StatusCode //状态码，0- 成功，其他值失败
	//2: optional string StatusMsg //返回状态描述
	//3: optional Comment Comment //评论成功返回评论内容，不需要重新拉取整个列表

	//生成回应结构体
	resp = new(userplat.CommentActionResponse)
	//校验参数
	err = req.IsValid()
	if err != nil {
		resp.StatusCode = 1203
		return resp, err
	}
	//实现逻辑
	//StatusCode, StatusMsg, Comment  resp
	resp, err = userplatservice.UserCommentActionService(ctx, req)
	if err != nil {
		resp.StatusCode = 1204
		return resp, err
	}

	return resp, nil

}

// UserCommentList implements the UserPlatServiceImpl interface.
func (s *UserPlatServiceImpl) UserCommentList(
	ctx context.Context,
	req *userplat.CommentListRequest,
) (resp *userplat.CommentListResponse, err error) {
	//1: i32 StatusCode //状态码，0- 成功，其他值失败
	//2: optional string StatusMsg //返回状态描述
	//3: list<Comment> CommentList //评论列表

	//生成回应结构体
	resp = new(userplat.CommentListResponse)
	//校验参数
	err = req.IsValid()
	if err != nil {
		resp.StatusCode = 1203
		return resp, err
	}
	//实现逻辑
	//StatusCode, StatusMsg, CommentList
	resp, err = userplatservice.UserCommentListService(ctx, req)
	if err != nil {
		resp.StatusCode = 1204
		return resp, err
	}

	return resp, nil
}
