package uuservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/useruser/dal/db"
	"github.com/808-not-found/tik_duck/cmd/useruser/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserRelationFollowListService(
	ctx context.Context,
	req *useruser.RelationFollowListRequest,
) (*useruser.RelationFollowListResponse, error) {
	var resp useruser.RelationFollowListResponse

	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 1001
		return &resp, err
	}
	// 验证登录状态
	myID := claims.ID
	if myID == 0 {
		resp.StatusCode = 1002
		return &resp, err
	}
	// 查询数据库
	var dbUsers []*db.User
	dbUsers, err = db.GetFollowList(myID)
	if err != nil {
		resp.StatusCode = 1006
		return &resp, err
	}
	// 数据封装
	rpcUsers, err := pack.Users(dbUsers, myID)
	if err != nil {
		resp.StatusCode = 1007
		return &resp, err
	}
	resp.UserList = rpcUsers

	return &resp, nil
}