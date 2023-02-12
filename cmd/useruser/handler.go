package main

import (
	"context"

	useruser "github.com/808-not-found/tik_duck/kitex_gen/useruser"
)

// UserUserServiceImpl implements the last service interface defined in the IDL.
type UserUserServiceImpl struct{}

// UserRelationAction implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationAction(
	ctx context.Context,
	req *useruser.RelationActionRequest,
) (resp *useruser.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRelationFollowList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFollowList(
	ctx context.Context,
	req *useruser.RelationFollowListRequest,
) (resp *useruser.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRelationFollowerList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFollowerList(
	ctx context.Context,
	req *useruser.RelationFollowerListRequest,
) (resp *useruser.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRelationFriendList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFriendList(
	ctx context.Context,
	req *useruser.RelationFriendListRequest,
) (resp *useruser.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
