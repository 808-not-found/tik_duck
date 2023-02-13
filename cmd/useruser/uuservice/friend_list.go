package uuservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/useruser/dal/db"
	"github.com/808-not-found/tik_duck/cmd/useruser/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserRelationFriendListService(
	ctx context.Context,
	req *useruser.RelationFriendListRequest,
) (*useruser.RelationFriendListResponse, error) {
	var resp useruser.RelationFriendListResponse
	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 1001
		return &resp, nil
	}
	// 检查登录状态
	myID := claims.ID
	if myID == 0 {
		resp.StatusCode = 1002
		return &resp, nil
	}
	// 查询数据库
	var dbUsers []*db.User
	dbUsers, err = db.GetFriendList(ctx, myID)
	if err != nil {
		resp.StatusCode = 1009
		return &resp, nil
	}
	// 封装数据
	var rpcUsers []*useruser.User
	rpcUsers, err = pack.Users(dbUsers, myID)
	if err != nil {
		resp.StatusCode = 1010
		return &resp, nil
	}
	resp.UserList = rpcUsers

	return &resp, nil
}