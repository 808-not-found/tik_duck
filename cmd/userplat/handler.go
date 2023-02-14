package main

import (
	"context"

	userplat "github.com/808-not-found/tik_duck/kitex_gen/userplat"
)

// UserPlatServiceImpl implements the last service interface defined in the IDL.
type UserPlatServiceImpl struct{}

// UserFavoriteAction implements the UserPlatServiceImpl interface.
func (s *UserPlatServiceImpl) UserFavoriteAction(
	ctx context.Context,
	req *userplat.FavoriteActionRequest,
) (resp *userplat.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// UserFavoriteList implements the UserPlatServiceImpl interface.
func (s *UserPlatServiceImpl) UserFavoriteList(
	ctx context.Context, req *userplat.FavoriteListRequest,
	) (resp *userplat.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// UserCommentAction implements the UserPlatServiceImpl interface.
func (s *UserPlatServiceImpl) UserCommentAction(
	ctx context.Context, req *userplat.CommentActionRequest,
	) (resp *userplat.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// UserCommentList implements the UserPlatServiceImpl interface.
func (s *UserPlatServiceImpl) UserCommentList(
	ctx context.Context, req *userplat.CommentListRequest,
	) (resp *userplat.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
