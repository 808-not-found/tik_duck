package uuservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/useruser/dal/db"
	"github.com/808-not-found/tik_duck/cmd/useruser/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserRelationFollowerList(
	ctx context.Context,
	req *useruser.RelationFollowerListRequest,
) (*useruser.RelationFollowerListResponse, error) {
	var resp useruser.RelationFollowerListResponse
	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 3009
		return &resp, err
	}
	// 登录状态检查
	myID := claims.ID
	if myID == 0 {
		resp.StatusCode = 3010
		return &resp, err
	}
	// 请求数据库
	var dbUsers []*db.User
	dbUsers, err = db.GetFollowerList(ctx, req.UserId)
	if err != nil {
		resp.StatusCode = 3011
		return &resp, err
	}
	// 封装数据
	rpcUsers, err := pack.Users(dbUsers, myID)
	if err != nil {
		resp.StatusCode = 3012
		return &resp, err
	}
	resp.UserList = rpcUsers
	return &resp, nil
}
